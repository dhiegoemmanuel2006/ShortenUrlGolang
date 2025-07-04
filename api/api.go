package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/api/shorten", HandlerPost(db))
	r.Get("/api/{code}", HandlerGet(db))
	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func HandlerPost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		}
		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid url passeed"}, http.StatusBadRequest)
		}
		code := genCode()
		db[code] = body.URL

		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characteres = "abcedefghiojklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ1234567890"

func genCode() string {
	byts := make([]byte, 8)
	const n = 8
	for i := range n {
		byts[i] = characteres[rand.IntN(len(characteres))]
	}

	return string(byts)
}

func HandlerGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		data, ok := db[code]
		if !ok {
			http.Error(w, "Url não encontrada", http.StatusNotFound)
		}
		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}

}
