# rps-backend

A funny microservice for playing *rock-paper-scissors*.

## Technical details

The app is written in [Go](https://go.dev/) using the [Gin framework](https://github.com/gin-gonic/gin).  

### CI/CD
On every git [tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging) push, the [Github Actions](https://github.com/features/actions) pipeline builds a new [Docker](https://www.docker.com/) image, 
pushes it to my [Registry](https://docs.docker.com/registry/), which starts a rolling update across the [Swarm](https://docs.docker.com/engine/swarm/) deployment.