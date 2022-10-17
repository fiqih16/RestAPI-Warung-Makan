## Dependency in this Go program

```sh
gorm.io/gorm
gorm.io/driver/mysql
github.com/gin-gonic/gin
github.com/joho/godotenv
github.com/dgrijalva/jwt-go
github.com/golang/crypto
github.com/mashingan/smapping
```

## Installation

- `git clone https://github.com/fiqih16/RestAPI-Warung-Makan`
- `cd RestAPI-Warung-Makan`
- Edit `.env`
  - Set your database connection details
- `go run main.go`

## API Endpoints

| Methods  | Endpoints            | Description                                                           |
| :------- | :------------------- | :-------------------------------------------------------------------- |
| `POST`   | api/auth/login       | login account must given `email` & `password` to body request         |
| `POST`   | api/auth/register    | Register account must given `name`,`email`,`password` to body request |
| `GET`    | api/customer/profile | Get Customer                                                          |
| `PUT`    | api/customer/profile | Edit Customer                                                         |
| `GET`    | api/menu/            | Get All Menu                                                          |
| `POST`   | api/menu/            | Create Menu                                                           |
| `GET`    | api/menu/:id         | Get Menu By Id                                                        |
| `PUT`    | api/menu/:id         | Edit Menu By Id                                                       |
| `DELETE` | api/menu/:id         | Delete Menu By Id                                                     |
| `POST`   | api/transaction/     | Create Transaction                                                    |
| `GET`    | /                    |                                                                       |
