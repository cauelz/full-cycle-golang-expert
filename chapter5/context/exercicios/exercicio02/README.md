# Exercício 2 — Contexto com cancelamento

## Conceitos
- O `context.WithCancel` permite criar um contexto que pode ser cancelado manualmente.
- Isso é útil para encerrar operações ou goroutines quando não são mais necessárias.

## Enunciado
Implemente um contexto com cancelamento usando `context.WithCancel` e cancele o contexto após 2 segundos.

> _Relembre o exercício 1 sobre criação de contextos._ 