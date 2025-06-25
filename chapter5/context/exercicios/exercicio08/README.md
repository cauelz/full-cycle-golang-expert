# Exercício 8 — Leitura concorrente com cancelamento

## Conceitos
- Leitura eficiente de arquivos grandes em Go
- Uso de goroutines e channels
- Cancelamento de operações com context

## Enunciado
Implemente um programa que leia um arquivo de texto grande linha a linha, usando uma goroutine para cada bloco de linhas (ex: 1000 linhas por goroutine). Use um channel para enviar as linhas lidas para o processamento principal. Utilize um `context.Context` para permitir o cancelamento da leitura a qualquer momento (por exemplo, após um tempo limite ou sinal do usuário).

> _Objetivos: Praticar leitura eficiente de arquivos grandes, uso de channels e cancelamento com context._ 