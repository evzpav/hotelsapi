# Hotels API


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
- 
- https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL