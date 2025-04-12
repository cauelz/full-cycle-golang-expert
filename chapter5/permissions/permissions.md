# Permissões

Em Go, as permissões de arquivos e diretórios são baseadas no modelo do Unix, que define três categorias de usuários e três tipos de permissões. Vamos detalhar esses conceitos:

---

### **1. Categorias de Usuários**
As permissões são definidas para três grupos:
1. **Owner (Dono)**: Usuário proprietário do arquivo/diretório.
2. **Group (Grupo)**: Usuários pertencentes a um grupo associado ao arquivo/diretório.
3. **Others (Outros)**: Todos os demais usuários do sistema.

---

### **2. Tipos de Permissões**
Cada categoria pode ter três tipos de acesso:
- **Read (Leitura)**: Representado por `r` (valor `4` em octal).
  - Arquivo: Permite ler o conteúdo.
  - Diretório: Permite listar os arquivos dentro dele.
- **Write (Escrita)**: Representado por `w` (valor `2` em octal).
  - Arquivo: Permite modificar o conteúdo.
  - Diretório: Permite criar, renomear ou excluir arquivos dentro dele.
- **Execute (Execução)**: Representado por `x` (valor `1` em octal).
  - Arquivo: Permite executar o arquivo como programa/script.
  - Diretório: Permite acessar o diretório (ex: usar `cd`).

---

### **3. Representação Octal**
As permissões são codificadas em um número octal de 3 dígitos, onde cada dígito corresponde a uma categoria (owner, group, others).  
Cada dígito é a soma dos valores das permissões:

| Permissão | Valor |
|-----------|-------|
| Read (r)  | 4     |
| Write (w) | 2     |
| Execute (x)| 1     |

**Exemplo**:
- `rwxr-xr--` é representado por `755` em octal:
  - Owner: `4+2+1 = 7` (rwx)
  - Group: `4+0+1 = 5` (r-x)
  - Others: `4+0+0 = 4` (r--)

---

### **4. Uso em Go**
No pacote `os`, as permissões são definidas usando o tipo `os.FileMode`, que é um número octal. Exemplos comuns:

| Permissão (Octal) | Significado                     |
|-------------------|---------------------------------|
| `0644`            | Owner: rw-, Group: r--, Others: r-- |
| `0755`            | Owner: rwx, Group: r-x, Others: r-x |
| `0600`            | Apenas o owner pode ler/escrever     |

**Exemplo de código**:
```go
// Cria um arquivo com permissões 0644
arquivo, err := os.Create("exemplo.txt")
if err != nil {
    log.Fatal(err)
}
defer arquivo.Close()

// Altera as permissões para 0755
err = os.Chmod("exemplo.txt", 0755)
if err != nil {
    log.Fatal(err)
}

// Cria um diretório com permissões 0755
err = os.Mkdir("meudir", 0755)
if err != nil {
    log.Fatal(err)
}
```

---

### **5. Observações Importantes**
- **Umask**: O Go respeita a máscara de permissões do sistema (`umask`). Por exemplo, se você definir `0666` (rw-rw-rw-), o arquivo pode ser criado com `0644` devido ao `umask` padrão (022), que remove permissões de escrita para group/others.
- **Diretórios**: A permissão de execução (`x`) é necessária para acessar o conteúdo do diretório.
- **Arquivos Executáveis**: Para scripts ou binários, use `0755` para permitir execução.

---

### **6. Constantes Úteis**
O pacote `os` define constantes para facilitar a combinação de permissões:
```go
os.ModePerm        // 0777: rwx para todos
os.ModeDir         // 040000: Identifica um diretório
os.ModeAppend      // Arquivo só pode ser aberto em modo append
os.ModeExclusive   // Arquivo é bloqueado para acesso exclusivo
```

---

### **7. Representação Simbólica**
Além do octal, você pode usar a notação simbólica (ex: `rwxr-xr--`), mas em Go é mais comum usar octal diretamente.

---

### Resumo Prático:
- Use `0644` para arquivos comuns (dono lê/escreve, outros só leem).
- Use `0755` para diretórios ou arquivos executáveis.
- Use `0600` para dados sensíveis (apenas o dono). 

Se precisar de mais detalhes ou exemplos específicos, é só perguntar! 😊