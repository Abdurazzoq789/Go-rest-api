Taths so simple go rest api Project 
============================

DIRECTORY STRUCTURE
-------------------

        app/        Bu joyda barja logikalar turadi
            database/   Database bilan bog'liq ishlar
            models/     Modellar
            test/   Oddiy test yozamiz

        main.go     Executable file

REQUIREMENTS
------------

Minimum Go version 1.1

INSTALLATION
------------
Birinchi navbatta config file to'g'irlanadi. Ya'ni  app/database/config.go fildagi configuratsialar to'girlanadi so'gra

```go
var (
    dbUsername = "your username"
    dbPassword = "your password"
    dbHost = "your host maybe: localhost"
    dbTable = "your table name"
    dbPort = "psql connected port"
    pgConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
    dbHost, dbPort, dbUsername, dbTable, dbPassword)
)
```

So'ngra

~~~
go run main.go
~~~

buyruq ishga tushiriladi va

~~~
http://localhost:9002/
~~~

End point orqali application ishga tushiriladi


AndPoints
------------

~~~
    GET http://localhost:9002/
    GET http://localhost:9002/api/posts
    POST http://localhost:9002/api/posts
    GET http://localhost:9002/api/post/{id}
    PUT http://localhost:9002/api/post/{id}
    DELETE http://localhost:9002/api/post/delete/{id}
~~~