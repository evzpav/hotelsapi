# Go Hotels API

Exemplo de API em Golang usando [Gorilla Mux](https://github.com/gorilla/mux) e [GORM](https://gorm.io/)

1) Criar rotas, handlers e converter json->struct struct->json:

    - Criar struct hotel
    ```json
        {
            "name": "Costão do Santinho",
            "city": "Florianópolis",
            "nrOfEmployees": 100,
            "revenue": 2342343.99,
            "active": true
        }

    ```
    - Pegar boiler plate de qualquer rota e jogar na main: https://github.com/gorilla/mux#examples
    - Criar GET, city como queryParam e retorna array de Hotels
    - Criar POST, recebe hotel, parse do body, unmarshal

2) Conectar Postgres e fazer insert e select no banco:
    - Adicionar lib para acesso ao Postgres: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
    - Inicializar Postgres (*gorm.DB)
    - Criar handler struct e passar *gorm.DB
    - Criar tabela no Postgres init.sql
    - Criar insert com GORM
    - Criar select com GORM
3) Refatorar em pacotes
4) Adicionar test unitários
