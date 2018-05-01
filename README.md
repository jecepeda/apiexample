# Welcome to APIExample!

This is an example of how to build a simple CRUD App with [Buffalo](https://gobuffalo.io)

## The goal

The goal of this project is to try how is to develop web applications with [Golang](https://golang.org/). Also is a simple project to test 
new frontend frameworks and having a basic API to do basic CRUDs.

## Requirements
- Docker.
- Docker-Compose.
- Go (1.10>=).

## Starting the Application

First, we need to configure the environment variables. To do so, create a new .env file with the following attributes:
- `POSTGRES_USER`: The postgresql user that you would like to use (e.g. postgres).
- `POSTGRES_PASSWORD`: The postgresql password that you would like to use (e.g. postgres).

Then, the first step is to create the database and run the application:
- Run `docker-compose run web buffalo db create`: it will build everything and will create the database. You can also use `make create`.
- Run `docker-compose up` to start the application. You can also use `make up`

## Extras

Feel free to fork and add new features to this application. If you want to test the application you can just type:
- `docker-compose run web buffalo test`

If you want to add more resources, actions, endpoints, etc. you can go to the main [Buffalo](https://gobuffalo.io) page and look at the docs
