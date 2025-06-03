# ENCURTADOR DE URL GOLANG

## 🚀 Como Rodar o Projeto

### ✅ Opção 1: Rodando com Docker (recomendado)

#### 1. Pré-requisitos

- Docker instalado: https://docs.docker.com/get-docker/

#### 2. Clonando o repositório

```bash
git clone https://github.com/dhiegoemmanuel2006/ShoirtenUrlGolang.git
cd ShoirtenUrlGolang
```
#### 3. Build da imagem
```bash
docker build -t shorten-url-golang .
```
#### 4. Rodando o container
````bash
docker run -p 8080:8080 shorten-url-golang
````
### ⚙️ Opção 2: Rodando Localmente (sem Docker)

#### 1. Verifique se o Go está instalado
```bash
go version
```
#### 2. Clonar o projeto
```bash
git clone https://github.com/dhiegoemmanuel2006/ShoirtenUrlGolang.git
cd ShoirtenUrlGolang

```
#### 3. Instalar as dependências
```bash
go mod tidy
```

#### 4. Executar a aplicação
```bash
go run .
```


### Requisição para pedir para criar a URL curta
#### localhost:8080/api/shorten
```json
{
  "url": "AQUI A URL QUE DESEJA ENCURTAR"
}
```
### Respostas que vai receber 
```json
{
  "data": "CÓDIGO DA URL QUE FOI ENCURTADA"
}
```
### Requisição com o código
#### localhost:8080/api/{código da url que pediu}
##### Após isso você direcionado para onde o código da URL estiver apontando dentro do sistema, se não houver. Então sera gerado um erro "NOT FOUND"