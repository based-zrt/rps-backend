# rps-backend

A funny microservice for playing *rock-paper-scissors*.

## Technical details

The app is written in [Go](https://go.dev/) using the [Gin framework](https://github.com/gin-gonic/gin).  

### CI/CD
On every git [tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging) push, the [GitHub Actions](https://github.com/features/actions) pipeline builds a new [Docker](https://www.docker.com/) image, 
pushes it to my [Registry](https://docs.docker.com/registry/), which starts a rolling update across the [Swarm](https://docs.docker.com/engine/swarm/) deployment.

### Development
Given Go's easy and light approach on build complexity, no specific environment needed,
however I strongly recommend [JetBrains GoLand](https://www.jetbrains.com/go/).

### Testing
Load testing is done with [k6](https://k6.io/).  
You can perform the test by running:
```shell
cd .k6
k6 run -e URL=<your host url> load_test.js
```