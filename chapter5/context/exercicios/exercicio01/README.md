# Exercício 1 — Criar um contexto vazio

## Conceitos
- O contexto em Go é utilizado para carregar prazos, cancelamentos e outros valores entre processos e goroutines.
- O ponto de partida para a maioria dos contextos é o `context.Background()`, que retorna um contexto vazio, geralmente usado em funções de nível superior.

## Enunciado
Escreva um programa que apenas cria um contexto vazio usando `context.Background()`.

> _Nota: Este é o ponto de partida para todos os exercícios seguintes._ 