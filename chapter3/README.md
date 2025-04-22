# Interfaces em Go

Este capítulo aborda um dos conceitos mais poderosos e flexíveis da linguagem Go: as Interfaces. Aqui você aprenderá como usar interfaces para criar código mais modular, testável e reutilizável.

## Índice
1. [Conceitos Básicos](docs/01-conceitos-basicos.md)
2. [Tipos de Interfaces](docs/02-tipos-de-interfaces.md)
3. [Implementação Implícita](docs/03-implementacao-implicita.md)
4. [Composição de Interfaces](docs/04-composicao-de-interfaces.md)
5. [Interfaces Vazias](docs/05-interfaces-vazias.md)
6. [Type Assertions e Type Switches](docs/06-type-assertions.md)
7. [Boas Práticas](docs/07-boas-praticas.md)

## Estrutura do Capítulo

### Documentação
A pasta `docs/` contém documentação detalhada sobre cada tópico relacionado a interfaces em Go.

### Exemplos
A pasta `exemplos/` contém implementações práticas dos conceitos apresentados na documentação.

### Exercícios
A pasta `exercicios/` contém desafios e exercícios para praticar os conceitos aprendidos.

## Recursos Adicionais

1. **Documentação Oficial**:
   - [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)
   - [Go by Example - Interfaces](https://gobyexample.com/interfaces)

2. **Interfaces Comuns**:
   - `io.Reader`
   - `io.Writer`
   - `fmt.Stringer`
   - `error`

3. **Dicas de Debug**:
   - Use `fmt.Printf("%T\n", x)` para ver o tipo real de uma interface
   - Use `reflect.TypeOf(x)` para inspeção mais detalhada 