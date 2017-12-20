![https://github.com/monstar-lab/goweb1/](https://pictr.com/images/2017/12/16/acd981ebabeb4f02458a0d96bcde53a0.jpg)


# goweb1 sample clean architechture

This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob

 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :

 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Handler Layer

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How To Run This Project

```bash
# GET WITH GO GET
go get github.com/monstar-lab/goweb1 

# Go to directory
cd $(go env GOPATH)/src/github.com/monstar-lab/goweb1

# Up containers
docker-compose up --build

```
### [>>Learn more about this docker environment](https://github.com/monstar-lab/goweb1/wiki/Setting-up-the-development-environment)

### About code samples