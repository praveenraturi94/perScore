# per score server
This service is a primary interface that provide Routing between services and act as a controller. It implements both REST and GRPC calls in order to interacts between clients and its internal services

### Dependencies

* [Postgresql](https://wiki.westfieldlabs.com/display/WL/PostgreSQL)
* [GORM](https://github.com/jinzhu/gorm)
* [Go](https://wiki.westfieldlabs.com/display/WL/Go)

### Build and run this project

1. To give privilege to ur .sh file
    ```
  chmod +x setupPostgres.sh
    ```
2. Run .sh file to create role and database
    ```
    ./setupPostgres.sh
    ```
3. Run command to migrate database
    ```
    go run main.go createDB
    ```
4. Run command to start server
    ```
    go run main.go serve
    ```
