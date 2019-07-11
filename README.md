# Sumc API
Sumc API is **high-available** service, wrapping Sofia Urban Mobility Center API and providing realtime schedules of Sofia public transport. It is build using internal job dispatcher - workers architecture along with caching which provide high availability and supports further scalability.
### Endpoints
```
GET:  /stop/:number/ - returns the schedule for given stop's number. 
```
### Building
Project is using ```dep``` https://github.com/golang/dep as package manager.
```
go get https://github.com/golang/dep - download the dep package
dep ensure - installing dependencies in vendor folder
```

### Adding new packages to the project
Dep executable could be found in bin/ folder of your GOPATH instalation directory.
When there is a need to add new library to the project - simply "go get" it and then run the following command from your project directory:
```
dep ensure
* dep executable is located in bin folder inside your GO path
```
Dep will go through your local GOPATH directory and add files there into Gopkg.lock, and into the vendor folder.

### To do
- [x] Initiate working code
- [ ] Implement rate limiter
- [ ] Implement interfaces & DI
- [ ] Unit tests
