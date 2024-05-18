# Dating Apps
This is a Dating Application write using Go (Golang) with Echo Framework.

## 1. Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### 1.1. Prerequisites

These are the prerequisites to run the project:

- [Go](https://golang.org/doc/install)
- [mockgen](https://github.com/uber-go/mock)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)
- [GNU Make](https://www.gnu.org/software/make/)

### 1.2. Installing

A step-by-step series of examples that tell you how to get a development env running

1. Clone the repository

```bash
git clone
```

2. Intialize the project

```bash
make init
```

3. Open the cloned repository. We recommend you to use Visual Studio Code because it's free and light but powerfull. [Visual Studio Code](https://code.visualstudio.com/download)

4. You need to configure `config.yaml.example` into a `.yaml` file and use it as your config template before running the project. Please note this config will not be uploaded to the repository.

5. This project has set to port `:9000`, but you can change to your desire port. To run the project.

```bash
go run main.go
```

6. Build the project
```bash
go build
```

## 2. Working Directory

This project also have many folder inside, then to minimize the complexity the directory will as follow:

```
project/
├── endpoint/
│   └── ...
├── helper/
│   ├── auth/
│   ├── database/
│   ├── static/
│   └── ...
├── http/
│   └── ...
├── initialization/
│   └── ...
├── model/
│   ├── base/
│   ├── entity/
│   ├── parameter/
│   ├── request/
│   ├── response/
│   └── ...
├── repository/
│   ├── user_repository/
│   └── ...
├── service/
│   ├── user_service/
│   └── ...
```

Explanations:

1. `endpoint` folder contains whole files for related the endpoint of the API.
2. `helper` folder contains whole files for help code easier such as auth logic, conversion logic, etc.
3. `http` folder contains files conversion result after processing and convert according with the response.
4. `initialization` folder contains initilization like server, database and etc.
5. `model` folder contains model entity represent of the database column.
6. `repository` folder contains files to querying the data to database or another data resource.
7. `service` folder contains files with business logic or other logics.

## 3. Usage of Go Library

This project use some libraries to support the developement process. And here are the libraries

1.  [faker](https://github.com/go-faker/faker) This library used to create fake or mock input data.
2.  [jwt](https://github.com/golang-jwt/jwt) This library used to generate and verify JWT token as an OAUTH security.
3.  [uuid](github.com/google/uuid) This library used to generate uuid.
4.  [echo](https://github.com/labstack/echo) This library used to as framework of creating the RestAPI.
5.  [cron](https://github.com/robfig/cron) This library used to run a cron job or scheduler.
6.  [viper](https://github.com/spf13/viper) This library used to help read config file.
7.  [testify](https://github.com/stretchr/testify) This library used to help testing process.
8.  [mock](https://github.com/uber-go/mock) This library used to create a mock file.
9.  [postgres](https://github.com/go-gorm/postgres) This library used to set the connection with PostgreSQL database.
10. [gorm](https://github.com/go-gorm/gorm) This library used to an ORM.# case-study-dealls
