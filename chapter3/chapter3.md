# Interfaces em Go

No Capítulo 2, exploramos como as Structs permitem criar estruturas de dados complexas para representar conceitos do mundo real de forma eficiente. Agora, com as Interfaces, damos um passo além: elas abstraem a implementação concreta das Structs, viabilizando princípios essenciais como `desacoplamento` e `polimorfismo`.

Vou deixar aqui um vídeo que eu gosto muito da balta.io [Orientação a objetos: Classe Abstrata VS Interface | por André Baltieri](https://www.youtube.com/watch?v=mWgeJdhrtDI) que explica muito bem a ideia por trás das abstrações de interfaces.

## Por que Interfaces são poderosas?

1- `Desacoplamento:` Interfaces definem o que um tipo deve fazer, sem impor como fazer. Por exemplo, em um sistema de pagamentos, você pode ter uma interface Pagador com o método ProcessarPagamento(), enquanto Structs como CartaoCredito ou Pix implementam a lógica específica.

```Go
type Pagador interface {
    ProcessarPagamento(valor float64) error
}
```

2- `Polimorfismo:` Uma função pode aceitar qualquer tipo que satisfaça uma interface. Seja um Usuario ou um Robo, desde que implementem `type Falante interface { Falar() string }`, ambos podem ser usados indistintamente.