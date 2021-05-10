# SWA -- Damn buggy login for a todo list web application

A simple web-based application, that allows two users to log in and manage their todos.

The two users are called 'Bob' and 'Doe'.

## Introduction

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Enjoy experimenting!

**Your help is required.**

Can you please harden our authentication mechanism. We have learned that we do not handle our cookies properly. Please check and repair anything you notice.

## Prerequisites

What things you need to have in the first place:

* [Golang](https://golang.org/) (Version 1.12 or above)
* [Gorilla Toolkit - Multiplexer](https://www.gorillatoolkit.org/pkg/mux)

## Installation

Follow these steps:

* The folder in which you unpacked the web app is denoted as "$TODOPATH" in the following (something like "/path/to/swa__prakt2_todo-02" for example).

* Enter the web application's directory

```CLI
cd $TODOPATH
```

* Build the application from the sources

```CLI
go build
```

* Start the server (you could also start the executable "swa__prakt2_todo-02" directly)

```CLI
go run todoServer.go
```

* Access the web application using your browser: [http://localhost:8080/](http://localhost:8080/)

## Built with

* [Golang](https://www.golang.org/) - The Go Programming Language

## Author

* **Luigi Lo Iacono**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
