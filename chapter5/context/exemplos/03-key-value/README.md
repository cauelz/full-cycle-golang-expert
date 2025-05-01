# Exemplo: Usando context com valores em Go

Este exemplo demonstra como utilizar o pacote `context` do Go para passar valores entre funções, utilizando o método `WithValue`.

## Código

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")

	bookHotel(ctx, "Caue")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
```

## Explicação

- Criamos um `context.Background()` como contexto base.
- Utilizamos `context.WithValue` para adicionar um valor ao contexto, neste caso, a chave `"token"` com o valor `"senha"`.
- Passamos o contexto para a função `bookHotel`, que pode acessar o valor do token através de `ctx.Value("token")`.
- O valor do token é impresso na tela.

## Saída esperada

```
senha
```

## Observações
- O uso de `context.WithValue` é recomendado apenas para dados que realmente fazem parte do contexto de execução, como tokens de autenticação, deadlines, etc.
- Para passar parâmetros entre funções, prefira argumentos explícitos sempre que possível. 