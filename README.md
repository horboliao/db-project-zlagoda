# db-project-zlagoda
project for university

Website for supermarket Zlagoda. Developed for managers and cashiers. 

Technologies:
- Gin, go
- React, js

You can:
• add/delete/create employees/customers/products/categories/products in store 
• see statistics about employees/customers/products/categories/products in store
• print reports of current tables
• sort tables using filters

To run:
• you need to instal go and docker
• docker compose up
• go mod tidy
• go get -u github.com/swaggo/files
• go get -u github.com/swaggo/gin-swagger
• cd api | swag init -g cmd/main.go
• go run ./... cmd
• cd web | npm start
