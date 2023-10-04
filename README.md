# rps-backend [![Build Status](https://github.com/based-zrt/rps-backend/actions/workflows/build.yml/badge.svg)](https://github.com/based-zrt/rps-backend/actions/workflows/build.yml)

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

### License
```
    rps-backend - Rock Paper Scissors backend
    Copyright (C) 2023  SunStorm

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
```