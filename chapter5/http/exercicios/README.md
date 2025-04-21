# Exercícios para praticar o pacote http

Este documento contém 50 exercícios para praticar os principais conceitos do pacote `http` em Go, organizados por nível de dificuldade e com foco na repetição de conceitos fundamentais.

## Exercícios Básicos

1. Crie um servidor HTTP simples que responda "Hello, World!" na rota raiz ("/").

2. Implemente um servidor que retorne a hora atual quando acessado.

3. Crie um handler que retorne diferentes mensagens baseadas no método HTTP (GET, POST, PUT, DELETE).

4. Desenvolva um servidor que conte e exiba o número de vezes que foi acessado.

5. Crie um endpoint que receba um nome como parâmetro de query e retorne "Olá, {nome}!".

6. Implemente um servidor que retorne diferentes status codes baseado em uma rota específica.

7. Crie um middleware que registre todas as requisições (método, rota e timestamp).

8. Desenvolva um endpoint que retorne um JSON simples com informações sobre o servidor.

9. Implemente um handler que processe dados de um formulário HTML simples.

10. Crie um servidor que sirva arquivos estáticos de um diretório específico.

11. Desenvolva um endpoint que aceite uploads de arquivos.

12. Implemente rate limiting básico usando um middleware.

13. Crie um handler que retorne headers personalizados na resposta.

14. Desenvolva um endpoint que demonstre o uso de cookies.

15. Implemente um servidor que redirecione certas rotas para outros endpoints.

16. Crie um servidor que responda com diferentes tipos de conteúdo (text/plain, application/json, text/html).

17. Desenvolva um endpoint que processe parâmetros de URL (/users/{id}).

18. Implemente um handler que leia e processe o corpo de uma requisição POST.

19. Crie um servidor que demonstre o uso de timeouts em requisições.

20. Desenvolva endpoints para demonstrar diferentes métodos de codificação (base64, URL encoding).

21. Implemente um servidor que utilize query strings para filtrar dados.

22. Crie um handler que demonstre o uso de context para cancelamento de requisições.

23. Desenvolva um endpoint que utilize template HTML para renderizar respostas.

24. Implemente um servidor que demonstre o uso de sessions com cookies.

25. Crie um middleware para validação de headers específicos.

## Exercícios Intermediários

26. Crie uma API RESTful simples para um CRUD de tarefas (sem banco de dados, use memória).

27. Implemente autenticação básica (Basic Auth) em um servidor HTTP.

28. Desenvolva um proxy reverso simples que encaminhe requisições para diferentes servidores.

29. Crie um servidor que implemente rate limiting por IP.

30. Implemente um sistema de cache em memória para respostas HTTP.

31. Desenvolva um middleware de compressão GZIP para as respostas.

32. Crie um servidor que suporte tanto HTTP quanto HTTPS.

33. Implemente um sistema de roteamento personalizado sem usar frameworks externos.

34. Desenvolva um middleware de CORS completo.

35. Crie um servidor que implemente long polling.

36. Implemente um sistema de sessões em memória.

37. Desenvolva um middleware de validação de JWT.

38. Crie um servidor que suporte uploads de arquivos grandes com progress tracking.

39. Implemente um sistema de rate limiting com diferentes limites por rota.

40. Desenvolva um middleware de timeout para requisições longas.

41. Crie um servidor que implemente server-sent events (SSE).

42. Implemente um sistema de websockets simples.

43. Desenvolva um sistema de cache com invalidação por tempo.

44. Crie uma API que implemente paginação de resultados.

45. Implemente um middleware para logging detalhado de erros.

46. Desenvolva um servidor que suporte múltiplos formatos de resposta (JSON, XML, YAML).

47. Crie um sistema de autenticação com refresh tokens.

48. Implemente um middleware para validação de requisições JSON.

49. Desenvolva um sistema de roteamento com middlewares específicos por rota.

50. Crie um servidor que implemente graceful shutdown.

## Dicas para Resolução

- Comece pelos exercícios básicos e avance gradualmente
- Use a documentação oficial do Go como referência
- Teste cada implementação com diferentes cenários
- Pratique o tratamento adequado de erros
- Considere aspectos de performance e segurança
- Documente seu código
- Escreva testes para suas implementações
- Reutilize conceitos aprendidos em exercícios anteriores
- Combine diferentes conceitos em uma única solução

## Recursos Úteis

- [Documentação oficial do pacote http](https://pkg.go.dev/net/http)
- [Go by Example - HTTP Servers](https://gobyexample.com/http-servers)
- [Go Web Examples](https://gowebexamples.com/)

