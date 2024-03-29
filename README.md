# GO 1.22

Nesse repositório, eu testei duas das novas features adicionadas ao pacote `net/http`:
- Wildcards
- Métodos da requisição no caminho ("METHOD route")

Para executar, garanta que você possui golang v1.22, no mínimo. Você pode executar
a API com `go run main.go`, ou buildando o executável: `go build`, após isso, execute
o executável criado na pasta raiz do projeto (dependendo do SO a forma de fazer isso varia).

Depois disso, você pode testar a API usando curl, Insomnia, Postman ou equivalentes.