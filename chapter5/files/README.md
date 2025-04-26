# Introdução à Manipulação de Arquivos e Diretórios em Go

Este diretório reúne materiais explicativos sobre os principais tópicos relacionados à manipulação de arquivos e diretórios na linguagem Go. Cada arquivo aborda um tema específico, facilitando o estudo e a consulta.

## Sumário dos Tópicos

- [Criando Arquivos](./docs/criando_arquivos.md): Aprenda como criar arquivos em Go, entender o funcionamento da função `os.Create` e como evitar sobrescritas indesejadas.
- [Abrindo e Manipulando Arquivos](./docs/abrindo_arquivos.md): Veja como abrir arquivos para leitura, escrita ou ambos, utilizando `os.Open` e `os.OpenFile`, além de entender o uso de flags e permissões.
- [Manipulando Diretórios](./docs/manipulando_diretorios.md): Descubra como criar, remover, listar e manipular diretórios, incluindo permissões e operações recursivas.
- [Leitura e Escrita Eficiente (bufio)](./docs/io_buffer.md): Entenda como utilizar buffers para otimizar operações de leitura e escrita em arquivos grandes com o pacote `bufio`.
- [Variáveis de Ambiente](./docs/variaveis_ambiente.md): Saiba como acessar, definir e remover variáveis de ambiente no sistema operacional usando Go.
- [Permissões](./docs/permissions.md): Compreenda o modelo de permissões de arquivos e diretórios no estilo Unix, como representá-las em Go e sua importância para segurança e controle de acesso.

Cada tópico pode fazer referência ao conceito de permissões, fundamental para garantir a segurança e o correto funcionamento das operações com arquivos e diretórios. Consulte o arquivo [Permissões](./docs/permissions.md) sempre que tiver dúvidas sobre como definir ou interpretar permissões em Go. 