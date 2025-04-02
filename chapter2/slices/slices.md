# Slices

Slices são uma das estruturas de dados mais utilizadas em Go, pois permitem que tenhamos uma "lista" de elementos com tamanho ajustado automaticamente, diferentemente dos arrays (que possuem tamanho fixo).
Essa característica nos permite criar sequências de valores sem precisar definir um tamanho prévio, além de possibilitar a adição de novos elementos posteriormente.
Devido a essa flexibilidade e funcionalidade dinâmica, os slices são muito mais comuns no dia a dia do desenvolvimento em Go.

## Como funcionam internamente? [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

Internamente, um slice não armazena dados diretamente, mas atua como uma referência para um array subjacente. Ele é composto por três elementos:

Ponteiro para o primeiro elemento do array referenciado.

Comprimento (length): número de elementos atualmente no slice.

Capacidade (capacity): número máximo de elementos que o slice pode armazenar sem realocar memória.

Essa estrutura torna os slices leves e flexíveis, pois operam sobre um array existente, mas com controle dinâmico de tamanho e capacidade.