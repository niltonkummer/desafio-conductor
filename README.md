# Desafio Conductor

## Dependencias

- Go 1.16 >
- docker + docker-compose

## Rodando Local

#### Compilando e rodando local com go:
```sh
git clone https://github.com/niltonkummer/desafio-conductor.git
cd desafio-conductor 
make run
```

A aplicação irá rodar no endereço [localhost:5000](http://localhost:5000/).

##### Rodando com docker
```sh
$ docker-compose up --build
```

### Host Teste
`warm-bastion-37111.herokuapp.com`

## Exemplo

`GET /contas/{id}/transacoes.pdf`

```sh
curl -sv --location --request GET 'https://warm-bastion-37111.herokuapp.com/conductor/v1/api/contas/c13d1ec5-3215-472c-b856-ed4a83ee5c4d/transacoes.pdf' \
--header 'Authorization:Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MTg4MDgzNzcsImV4cCI6MTY1MDM0NDM3NywiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.hf6ChYTnOw4dKuK51SZQA20k0J1eFOBLRcC-wD6Xhk4' \
-o fatura.pdf && open fatura.pdf
```

## Documentação API

- [Swagger](https://warm-bastion-37111.herokuapp.com/swagger/)
