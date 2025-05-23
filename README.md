# Load Tester CLI

## 📖 O que é este projeto?

Este projeto é uma aplicação **CLI escrita em Go** para realizar **testes de carga** (stress test) em serviços web. Ele permite simular múltiplos requests simultâneos a uma URL, avaliando o desempenho e a estabilidade do serviço.

## ❓ O que é um Stress Test?

Um **stress test** é uma técnica utilizada para avaliar a capacidade de um sistema web ao ser submetido a uma alta carga de usuários ou requisições simultâneas. O objetivo é identificar:

- Limites de desempenho.
- Possíveis falhas ou gargalos.
- Estabilidade sob condições extremas.


## 🚀 Como usar

### ✅ Pré-requisitos

- Docker instalado.

### 🛠️ Passos

1. **Construa a imagem Docker:**

```bash
docker build -t load-tester .
```

2. **Execute o teste de carga:**

```bash
docker run load-tester ./load-tester --url=http://google.com --requests=1000 --concurrency=10
```

### ⚙️ Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de chamadas simultâneas.

## 📊 Relatório Gerado

Ao final da execução, será exibido um relatório contendo:

- **⏱️ Tempo total de execução:** duração completa do teste.
- **📥 Total de requests realizados:** quantidade de requisições enviadas.
- **✅ Quantidade de status HTTP 200:** sucesso nas respostas.
- **📊 Distribuição de outros códigos de status:** como erros `404`, `500` etc.

Esse relatório ajuda a entender como o serviço responde sob carga e quais melhorias podem ser feitas.



