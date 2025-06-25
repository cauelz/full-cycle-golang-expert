# Exercícios de Context em Go

Pratique o conceito de `context` em Go com os exercícios abaixo. Os exercícios variam do básico ao intermediário e são ideais para fixar o uso de contextos em aplicações Go.

## Exercícios

1. **Criar um contexto vazio:**
   - Escreva um programa que apenas cria um contexto vazio usando `context.Background()`.
   - _Nota: Este é o ponto de partida para todos os exercícios seguintes._

2. **Contexto com cancelamento:**
   - _Relembre o exercício 5 sobre criação de contextos._
   - Implemente um contexto com cancelamento usando `context.WithCancel` e cancele o contexto após 2 segundos.

3. **Contexto com timeout:**
   - _Relembre os exercícios 5 e 6 sobre criação e cancelamento de contextos._
   - Crie um contexto com timeout de 1 segundo usando `context.WithTimeout` e exiba uma mensagem quando o timeout ocorrer.

4. **Contexto com deadline:**
   - _Relembre o exercício 7 sobre timeout em contextos._
   - Utilize `context.WithDeadline` para criar um contexto que expira em 500ms e mostre quando ele for cancelado.

5. **Propagação de cancelamento:**
   - _Relembre os exercícios 6, 7 e 8 sobre cancelamento, timeout e deadline._
   - Crie uma função que recebe um contexto e cancele a execução de uma goroutine quando o contexto for cancelado.

6. **Passando valores no contexto:**
    - _Relembre o exercício 5 sobre criação de contextos._
    - Use `context.WithValue` para passar um valor (ex: userID) entre funções.

7. **Recuperando valores do contexto:**
    - _Relembre o exercício 10 sobre passagem de valores no contexto._
    - Implemente uma função que recupera um valor específico do contexto.

8. **Leitura concorrente com cancelamento:**
   - Implemente um programa que leia um arquivo de texto grande linha a linha, usando uma goroutine para cada bloco de linhas (ex: 1000 linhas por goroutine). Use um channel para enviar as linhas lidas para o processamento principal. Utilize um `context.Context` para permitir o cancelamento da leitura a qualquer momento (por exemplo, após um tempo limite ou sinal do usuário).
   - _Objetivos: Praticar leitura eficiente de arquivos grandes, uso de channels e cancelamento com context._

9. **Contagem de palavras com timeout:**
   - Crie um programa que conte o número de palavras em um arquivo grande, processando o arquivo em paralelo (divida o arquivo em partes e processe cada parte em uma goroutine). O programa deve aceitar um timeout via context: se o tempo acabar, o programa deve cancelar todas as goroutines e retornar o resultado parcial.
   - _Objetivos: Praticar uso de context com timeout, divisão de tarefas entre goroutines e coleta de resultados via channel._

10. **Busca de padrão com contexto de cancelamento:**
   - Implemente um programa que busque por uma palavra/padrão em um arquivo grande, usando múltiplas goroutines para processar diferentes partes do arquivo. O programa deve parar imediatamente todas as buscas assim que encontrar a primeira ocorrência, usando context para cancelar as goroutines restantes.
   - _Objetivos: Praticar busca concorrente em arquivos, uso de context para cancelamento imediato e sincronização de goroutines com channels._

11. **Pipeline de processamento com context:**
   - Monte um pipeline de processamento de linhas de um arquivo:
     1. Uma goroutine lê as linhas do arquivo e envia para um channel.
     2. Uma ou mais goroutines processam as linhas (ex: transformam o texto, filtram, etc) e enviam para outro channel.
     3. Uma goroutine final coleta e salva o resultado.
   - Implemente o cancelamento do pipeline usando context.
   - _Objetivos: Praticar construção de pipelines com channels, uso de context para cancelar todo o pipeline e manipulação eficiente de arquivos grandes._

12. **Cancelando múltiplas goroutines:**
    - _Relembre os exercícios 6 e 9 sobre cancelamento e propagação em goroutines._
    - Lance 3 goroutines que escutam o contexto e finalize todas ao cancelar o contexto.

13. **Timeout em requisições HTTP:**
    - _Relembre o exercício 7 sobre timeout em contextos._
    - Faça uma requisição HTTP com timeout usando contexto.

14. **Contexto em banco de dados:**
    - _Relembre os exercícios 6, 7 e 9 sobre cancelamento, timeout e propagação._
    - Simule uma operação de banco de dados que pode ser cancelada via contexto.

15. **Hierarquia de contextos:**
    - _Relembre os exercícios 6, 7, 8 e 9 sobre criação e cancelamento de contextos._
    - Crie um contexto pai e dois filhos, cancele o pai e observe o efeito nos filhos.

16. **Contexto em servidor HTTP:**
    - _Relembre os exercícios 6, 7, 8, 9 e 13 sobre cancelamento, timeout e uso em HTTP._
    - Implemente um handler HTTP que respeita o cancelamento do contexto da requisição.

17. **Contexto em worker pool:**
    - _Relembre os exercícios 9 e 12 sobre goroutines e cancelamento._
    - Implemente um worker pool onde os workers param ao cancelar o contexto.

18. **Contexto em pipelines:**
    - _Relembre os exercícios 9, 12 e 17 sobre goroutines, cancelamento e worker pool._
    - Crie um pipeline de processamento de dados que pode ser interrompido via contexto.

19. **Contexto em testes:**
    - _Relembre os exercícios 7, 8 e 18 sobre timeout, deadline e pipelines._
    - Escreva um teste unitário que utiliza contexto com timeout para evitar deadlocks.

20. **Contexto customizado:**
    - _Relembre o exercício 5 sobre a interface básica de contextos._
    - Implemente um tipo customizado que implementa a interface `context.Context`.

21. **Contexto e select:**
    - _Relembre os exercícios 6, 7, 8 e 9 sobre cancelamento e uso de canais._
    - Use um `select` para aguardar entre um canal de dados e o canal de cancelamento do contexto.

22. **Contexto e logs:**
    - _Relembre os exercícios 10 e 11 sobre passagem e recuperação de valores no contexto._
    - Adicione informações do contexto (ex: requestID) nos logs de uma aplicação.

23. **Contexto e middlewares:**
    - _Relembre os exercícios 10, 11, 16 e 22 sobre valores no contexto, HTTP e logs._
    - Implemente um middleware HTTP que adiciona valores ao contexto da requisição.

24. **Contexto e gRPC:**
    - _Relembre os exercícios 7, 8, 9, 13 e 16 sobre timeout, cancelamento e uso em HTTP._
    - Simule uma chamada gRPC que utiliza contexto para deadline e cancelamento.

---

Bons estudos! Para cada exercício, tente implementar sozinho antes de buscar a solução. Se precisar de exemplos, consulte a documentação oficial: https://pkg.go.dev/context 

---

**Observação:** Agora cada exercício possui uma subpasta própria dentro de `exercicios/`, contendo um `README.md` com introdução ao conceito, dicas e o enunciado correspondente. Explore cada pasta para praticar e revisar os conceitos individualmente. 