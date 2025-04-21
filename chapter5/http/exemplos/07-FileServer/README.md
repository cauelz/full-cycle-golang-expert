# FileServer em Go

Este exemplo demonstra como criar um servidor de arquivos estáticos (FileServer) em Go, uma funcionalidade essencial para servir conteúdo web como HTML, CSS, JavaScript, imagens e outros recursos.

## O que é um FileServer?

Um FileServer é um servidor HTTP especializado em servir arquivos estáticos diretamente do sistema de arquivos. Em Go, o pacote `net/http` fornece o `http.FileServer`, que implementa essa funcionalidade de forma segura e eficiente.

## Características Principais

- **Streaming de Arquivos**: Utiliza streams para transferir arquivos, sendo eficiente com memória
- **Cache-Control**: Gerencia headers HTTP para cache automaticamente
- **Content-Type**: Detecta e define o tipo MIME correto dos arquivos
- **Range Requests**: Suporta download parcial de arquivos (útil para streaming de mídia)
- **Directory Listing**: Pode listar conteúdo de diretórios (configurável)
- **Segurança**: Previne directory traversal attacks

## Quando Usar

1. **Aplicações Web**: Servir arquivos HTML, CSS, JS e imagens
2. **APIs com Documentação**: Hospedar documentação estática
3. **Download de Arquivos**: Disponibilizar arquivos para download
4. **Mídia Streaming**: Servir arquivos de áudio/vídeo
5. **Single Page Applications (SPAs)**: Servir a aplicação frontend
