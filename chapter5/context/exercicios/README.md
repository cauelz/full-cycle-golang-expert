# Exercícios de Context em Go

Pratique o conceito de `context` em Go com os exercícios abaixo. Os exercícios variam do básico ao intermediário e são ideais para fixar o uso de contextos em aplicações Go.

## Exercícios

1. **Criar um contexto vazio:**
   - Escreva um programa que apenas cria um contexto vazio usando `context.Background()`.

2. **Contexto com cancelamento:**
   - Implemente um contexto com cancelamento usando `context.WithCancel` e cancele o contexto após 2 segundos.

3. **Contexto com timeout:**
   - Crie um contexto com timeout de 1 segundo usando `context.WithTimeout` e exiba uma mensagem quando o timeout ocorrer.

4. **Contexto com deadline:**
   - Utilize `context.WithDeadline` para criar um contexto que expira em 500ms e mostre quando ele for cancelado.

5. **Propagação de cancelamento:**
   - Crie uma função que recebe um contexto e cancele a execução de uma goroutine quando o contexto for cancelado.

6. **Passando valores no contexto:**
   - Use `context.WithValue` para passar um valor (ex: userID) entre funções.

7. **Recuperando valores do contexto:**
   - Implemente uma função que recupera um valor específico do contexto.

8. **Cancelando múltiplas goroutines:**
   - Lance 3 goroutines que escutam o contexto e finalize todas ao cancelar o contexto.

9. **Timeout em requisições HTTP:**
   - Faça uma requisição HTTP com timeout usando contexto.

10. **Contexto em banco de dados:**
    - Simule uma operação de banco de dados que pode ser cancelada via contexto.

11. **Hierarquia de contextos:**
    - Crie um contexto pai e dois filhos, cancele o pai e observe o efeito nos filhos.

12. **Contexto em servidor HTTP:**
    - Implemente um handler HTTP que respeita o cancelamento do contexto da requisição.

13. **Contexto em worker pool:**
    - Implemente um worker pool onde os workers param ao cancelar o contexto.

14. **Contexto em pipelines:**
    - Crie um pipeline de processamento de dados que pode ser interrompido via contexto.

15. **Contexto em testes:**
    - Escreva um teste unitário que utiliza contexto com timeout para evitar deadlocks.

16. **Contexto customizado:**
    - Implemente um tipo customizado que implementa a interface `context.Context`.

17. **Contexto e select:**
    - Use um `select` para aguardar entre um canal de dados e o canal de cancelamento do contexto.

18. **Contexto e logs:**
    - Adicione informações do contexto (ex: requestID) nos logs de uma aplicação.

19. **Contexto e middlewares:**
    - Implemente um middleware HTTP que adiciona valores ao contexto da requisição.

20. **Contexto e gRPC:**
    - Simule uma chamada gRPC que utiliza contexto para deadline e cancelamento.

---

Bons estudos! Para cada exercício, tente implementar sozinho antes de buscar a solução. Se precisar de exemplos, consulte a documentação oficial: https://pkg.go.dev/context 