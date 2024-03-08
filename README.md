# tech-challenge
Repositório destinado ao projeto da FIAP

## Rodando a aplicação com Kubernetes

### Garanta que você tenha o comando kubectl na sua máquina

#### Primeiro faça o build da sua imagem docker, utilize o comando abaixo:
```docker build . -t tech-challenge-go -f Dockerfile```

#### Após isso utilize o comando a baixo
```kubectl apply -f infra/```

#### Pronto, projeto rodando localmente em sua máquina com kubernetes


## Testes de estresse
### Tenha o K6 instalado em sua máquina e utilize o comando abaixo
```k6 run --duration 1m stress/stress.js```

# Documentação

Ao importar a documentação presente em `docs/tech-challenge.json` no Postman, terão valores de exemplos editáveis.
