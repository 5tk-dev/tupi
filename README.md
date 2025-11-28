# tupi

Um framework de validação e parsing de schemas em Go

- Geração de Fielder[T] para qualquer tipo (struct, map, slice, primitivos).
- Decode direto de map[string]string, JSON ou YAML.
- Validação com tags (required, min=18, max=99, etc).
- Regras customizáveis via tupi.SetRule.

---

Instalação

```sh
go get 5tk.dev/tupi
```

---

Uso Básico

```go
package main

import (
    "fmt"
    "5tk.dev/tupi"
)

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=18"`
}

func main() {
    // Parse gera o Fielder para o tipo
    // f := tupi.Parse[User]() -> retorna um fielder[User]
    f := tupi.Parse(&User{Name:"claudir"}) //retorn um Fielder[User] com default value

    // Decode a partir de um map
    schema := f.Decode(map[string]string{
        "Name":  "John",
        "Email": "john@example.com",
        "Age":   "17",
    })

    if schema.HasError() {
        fmt.Println("Erro:", schema.Error())
    } else {
        fmt.Println("Usuário válido:", schema.Value())
    }
}
```

---

Regras Customizadas

```go
// Registrar uma regra custom
tupi.SetRule("sanitize", func(v reflect.Value) bool {
    // exemplo simples: remover espaços
    if v.Kind() == reflect.String {
        return strings.TrimSpace(v.String()) != ""
    }
    return true
}, true) // true = executa antes de setar o valor
```

Uso na tag:

```go
type Post struct {
    Title string `validate:"required,sanitize"`
}
```

---

Features

- ✅ Cache de Fielder por tipo (registry interno com sync.Map).
- ✅ Decode de map, json, yaml.
- ✅ Validação com tags padrão e custom.
- ✅ Extensível com funções próprias.
- ✅ Imutável após criação → seguro para concorrência.

---

Roadmap

- [ ] Mais validadores built-in (regex, min/max length, etc).
- [ ] Melhor suporte a internacionalização de erros.
- [ ] Benchmarks comparativos com go-playground/validator.

---

Licença

MIT