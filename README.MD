<h1 align="center">RantauJoeang App - Backend</h1>

This API is created by me for build my personal project about rooming house e-commerce that will be helpful in future. This backend api using Golang with Gin and GORM package. [More about Golang](https://go.dev/)

## Built With

[![Go/Golang](https://img.shields.io/badge/Golang-1.20-cyan.svg?style=rounded-square)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-v.1.10-blue.svg?style=rounded-square)](https://gin-gonic.com/docs/)
[![GORM](https://img.shields.io/badge/Gorm-v.1.25-blue.svg?style=rounded-square)](https://gorm.io/)


## Requirements

1. <a href="https://go.dev/">Go</a>
2. <a href="https://gin-gonic.com/docs/">Gin</a>
3. <a href="https://gorm.io/">GORM</a>
4. <a href="https://www.mysql.com/">MySQL</a>
5. <a href="https://github.com/air-verse/air">Air</a>

## How to run the app ?

1. Clone this project
2. Open app's directory in CMD or Terminal
3. Turn on Web Server and MySQL can using Third-party tool like xampp, etc.
4. Create a database with the name #nama_database, and Import file sql to **phpmyadmin**
5. Open Postman desktop application or Chrome web app extension that has installed before
6. Choose HTTP Method and enter request url.(ex. localhost:3600/)
7. You can see all the end point [here](https://documenter.getpostman.com/view/14780095/2sAYdmmoQk)
8. Type `air server --port 8080` to activated the server (You can change 8080 to app port that you use in that project).

## Set up project

1. Create a file ".env" then setup environment variable like code below:

```
APP_PORT=<GIN LISTENING PORT>
DATABASE_HOSTNAME=<YOUR_DB_HOSTNAME>
DATABASE_PORT=<YOUR_DB_PORT>
DATABASE_USERNAME=<YOUR_DB_USERNAME>
DATABASE_PASSWORD=<YOUR_DB_PASSWORD>
DATABASE_NAME=<YOUR_DATABASE_NAME>
```

## Feature

1. Login and Register User
2. Admin
    1. Users API
        1. Get Users
        2. Get User
        3. Create User
        4. Update User
        5. Delete User
    2. Dorms API
        1. Get Dorms
        2. Get Dorm
        3. Create Dorm
        4. Update Dorm
        5. Delete Dorm
    3. Payments API
        1. Get Payments
        2. Get Payment
        4. Update Payment
    4. Invoices API
        1. Get Invoices
        2. Get Invoice
        3. Update Invoice
    5. Messages API
        1. Get Messages
        2. Get Message
        3. Create Message
        4. Reply Message
        5. Update Message
        6. Delete Message
3. Users
    1. Users API
        1. Get User
        2. Update User
    2. Dorms API
        1. Get Dorms
        2. Get Dorm
    3. Payments API
        1. Get Payments
        2. Get Payment
        3. Create Payment
    4. Invoices API
        1. Get Invoices
        2. Get Invoice
        3. Create Invoice
    5. Messages API
        1. Get Messages
        2. Get Message
        3. Create Message
        4. Reply Message
        5. Update Message
        6. Delete Message

## License

© [Muhammad Akbar Saladin Siregar](https://github.com/akbarsaladin36/)
