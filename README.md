# Desafio Go - Cliente e Servidor HTTP

Neste desafio, desenvolvemos um cliente e um servidor HTTP em Go para obter a cotação do dólar em tempo real e persisti-la em um banco de dados SQLite. O cliente solicita a cotação ao servidor, que por sua vez consome uma API externa, salva a cotação no banco de dados e retorna o resultado para o cliente.

## Requisitos do Desafio

- Implementar dois sistemas em Go: `client.go` e `server.go`.
- O `client.go` deve realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.
- O `server.go` deve consumir a API de câmbio de dólar e real no endpoint `https://economia.awesomeapi.com.br/json/last/USD-BRL` e retornar o resultado no formato JSON para o cliente.
- Utilizar o package `context` para definir timeouts.
- O `server.go` deve registrar cada cotação recebida no banco de dados SQLite.
- Timeout máximo de 200ms para chamar a API de cotação do dólar e de 10ms para persistir os dados no banco de dados.
- O `client.go` deve receber do `server.go` apenas o valor atual do câmbio (campo "bid" do JSON).
- Timeout máximo de 300ms para receber o resultado do `server.go`.
- Salvar a cotação atual em um arquivo "cotacao.txt".

## Estrutura do Projeto
- `client.go`: Implementação do cliente HTTP.
- `server.go`: Implementação do servidor HTTP.

## Contribuindo

Sinta-se à vontade para contribuir com melhorias, correções de bugs ou novos recursos. Basta abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
