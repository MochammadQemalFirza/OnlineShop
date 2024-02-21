# Online Shop Project

A Golang RESTful API with Postgresql project for online shop

# Table of Contents

* [Design](#design)
  * [REST API Design](#rest-api-design)
  * [ERD Design](#erd-design)
* [Features](#features)
* [Getting Started](#getting-started)
  * [Installation](#installation)
  * [Prequisited](#prequisited)
* [API Documentation](#api-documentation)

## Design 
### REST API Design
![Untitled Diagram drawio (2)](https://github.com/MochammadQemalFirza/OnlineShop/assets/90755886/26b3d058-96f4-4aa0-9425-a1f7f6423739)

### ERD Design
![Untitled Diagram drawio (6)](https://github.com/MochammadQemalFirza/OnlineShop/assets/90755886/c4106053-597f-46c2-b683-cf5794a760cc)

## Features
* Customer can view product list by product category
* Customer can add product to shopping cart
* Customers can see a list of products that have been added to the shopping cart
* Customer can delete product list in shopping cart
* Customers can checkout and make payment transactions
* Login and register customers

## Getting Started
### Installation
* Clone this repository
```
https://github.com/MochammadQemalFirza/OnlineShop.git
```
* Download this database file and open it on PgAdmin 
```
https://drive.google.com/drive/folders/1oktLw4O6YcIT-lRKYyCqf2jwSJILs4Rj?usp=sharing
```
* Install Library using terminal on your code editor
```
go get
```
* Setting .env with this template
```
#Database
DB_HOSTNAME="localhost"
DB_PORT="5432"
DB_USER="postgres"

DB_PASSWORD="your_paassword"

DB_DATABASE="olshop"
DB_SSLMODE="disable"

# Port
PORT = "8082"
```
* Run the server by typing this command on root folder project
```
go run server.go
```
### Prequisited
* [Postgresql](https://www.postgresql.org/download/) - SQL Database.
* [Golang](https://go.dev/dl/) - Programming Language


## API Documentation
* Customer can view product list by product category
  - method : GET
  - QueryParam : Key (Category) value (Makanan, Minuman, Pakaian)
  - API Endpoint : http://localhost:8082/product
  ![image](https://github.com/MochammadQemalFirza/OnlineShop/assets/90755886/3bdaa516-36d7-4fb6-897b-4eecc2f3993f)

