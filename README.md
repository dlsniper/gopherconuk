# Go Web Templates

This is a Go Web Server Swiss Army Knife, where various web tools are explored and prototyped in Golang, and referenced to quickly get basic boilerplate projects going.

## About

The base of this project was built off of [this one](https://github.com/dlsniper/gopherconuk) by Florin Pățan and I suggest you have a look at it for useful links and resources:

- It shows how to use [net/http](https://godoc.org/net/http), and how to structure a Go project.

- Dependency injection is used to insert a logger instance into the handler.

- You can also notice how the test is constructed in order to provide testing for the handler.

- A Docker container is available, thanks to the Dockerfile. It shows how to construct such containers.

On top of that the project includes the following:

- A mysql database connection - using [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql#usage) through [SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)

- A basic websocket endpoint using (Gorilla WebSocket)[github.com/gorilla/websocket]

## Usage

Clone this anywhere onto your computer and open the project in your editor.

If you prefer to use Go Mod run the following (using your own workspace directory):

```
export GO111MODULE=on && go mod init github.com/PeloDev/go-server-template
```

You can also create a `.env` file based on the [.env.example template](.env.example).

You will then have to export the variables before running the project. For example:

```
set -o allexport; source .env; set +o allexport
```

To run the project, in the root directory use:

```
go run main.go
```

Docker instructions pending...

## License

This project is under the MIT license. Please see the [LICENSE](LICENSE) file for more details.
