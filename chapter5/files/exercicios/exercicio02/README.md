# Exercício 02: Ler e Exibir o Conteúdo de um Arquivo

Escreva um programa em Go que:

1. Leia o conteúdo do arquivo `mensagem.txt` criado no exercício anterior.
2. Exiba o conteúdo lido no terminal.
3. Trate possíveis erros de leitura.

Dicas:
- Use as funções do pacote `os` e `io/ioutil` ou `os` e `io`. 

---

## Bônus: Por que fechar o arquivo antes de ler?

Quando você escreve em um arquivo, os dados geralmente não são gravados imediatamente no disco. Eles ficam em uma área de memória chamada **buffer**. O sistema operacional faz isso para otimizar a performance, agrupando várias operações de escrita antes de realmente gravar no disco.

- Enquanto o arquivo está aberto para escrita, parte dos dados pode estar apenas no buffer, e não no arquivo físico.
- Fechar o arquivo (`file.Close()`) força o sistema a gravar todos os dados pendentes no disco e libera o recurso do arquivo.

Se você tentar ler o arquivo antes de fechá-lo:
- Pode acabar lendo um arquivo incompleto, pois parte dos dados ainda não foi gravada no disco.
- Pode ocorrer conflito de acesso, dependendo do sistema operacional e do modo de abertura.

**Resumo:** Sempre feche o arquivo de escrita antes de abri-lo para leitura! Isso garante que tudo o que você escreveu foi realmente salvo no disco e estará disponível para leitura. 