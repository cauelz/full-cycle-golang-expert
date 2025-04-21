### O que é um *multiplexer* (ou “mux”) em Go?

Em aplicações web você tem **uma única porta** de entrada (o servidor) mas deseja encaminhar **várias rotas** para funções diferentes.  
O papel do *multiplexer* é exatamente esse: **examinar cada requisição HTTP e decidir qual *handler* deve atendê‑la**.  

Em Go esse roteador vem embutido no pacote `net/http` com o tipo `*http.ServeMux`, muitas vezes chamado só de “mux”. Ele é seguro para uso concorrente, não cria novas goroutines (isso fica a cargo do servidor) e usa regras bem claras de correspondência de padrões. citeturn0search0  

---

### Como o `ServeMux` decide o *handler*

1. **Correspondência literal do caminho** – Ex.: `"/about"` só atende exatamente `/about`.  
2. **Sub‑árvore** – Se o padrão termina em `/`, ele pega qualquer coisa abaixo. Ex.: `"/static/"` cobre `/static/css/main.css`.  
3. **Prioridade** – O padrão mais específico (mais longo) vence; se empatar, vale o primeiro registrado.  
4. **Padrões com host** – Você pode usar `"example.com/"` para diferenciar vários domínios no mesmo servidor.  

#### Novidades do Go 1.22+

A partir do Go 1.22 o roteador foi turbinado:

* **Correspondência por método** – Inclua o verbo: `"GET /users"` ou `"POST /users"`.  
* **Curingas com variáveis** – Use colchetes: `"GET /users/{id}"` e depois leia `r.PathValue("id")`.  
Essas melhorias tornam o `ServeMux` uma alternativa viável a outros roteadores externos. citeturn0search1turn0search3  

---

### Exemplo completo e enxuto

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// 1. Padrão literal
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, mundo!")
	})

	// 2. Sub‑árvore (arquivos estáticos)
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 3. Usando as novidades do Go 1.22
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")        // <- variável capturada no padrão
		json.NewEncoder(w).Encode(struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{ID: id, Name: "Usuário " + id})
	})

	// 4. Inicia o servidor
	http.ListenAndServe(":8080", mux)
}
```

*Dica:* se você não especificar um `ServeMux`, o pacote usa o `http.DefaultServeMux`, mas é boa prática criar o seu para evitar conflitos em aplicações maiores.

---

### Desafio prático 💪

1. **Objetivo:** Crie um pequeno serviço de tarefas (“TODO”).  
2. **Requisitos mínimos:**  
   - Rota `GET /todos` → devolve um *slice* de tarefas em JSON.  
   - Rota `POST /todos/{title}` → adiciona a tarefa em memória (não use banco por enquanto).  
   - Use o roteamento por método e variáveis do Go 1.22.  
3. **Extras (opcional):**  
   - Adicione uma rota `DELETE /todos/{id}`.  
   - Implemente *middleware* simples para registrar tempo de execução de cada requisição.

Experimente, depois rode `curl http://localhost:8080/todos` ou qualquer cliente HTTP e veja o mux trabalhando!  

Bons estudos e mãos à obra!