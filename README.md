1. Command migration:

    -  bee migrate -driver="postgres" -conn="postgres://postgres:123456@localhost:5432/goweb?sslmode=disable"

    -  bee migrate rollback -driver="postgres" -conn="postgres://postgres:123456@localhost:5432/goweb?sslmode=disable"

2. bee generate
    - bee generate controller hello