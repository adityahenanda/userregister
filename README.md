# userregister


clean architecture:

Models : layer for register model struct / request / response / entities
Repository : layer for database handler / query
Usecase : layer for bussiness logic
Delivery : output hanlder layer, http rquest / communication services
quick set up (clone to your gopath )

edit config.json (set your db connection), automatic migrate table
go get
go run main.go (running on localhost:7777)
