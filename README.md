Claerance
---------

This project intends to be a user-friendly, easy to set up authentication
"middleware" for [Traefik](https://github.com/containous/traefik).

## How-to's

### Docker
#### Building and running
Build: `docker build . -t claerance`  
Run: `docker run claerance`

### Not Docker
#### Building, testing & running
- Build: `go build -o claerance ./cmd/main.go`
- Test: `go test`
- Run: `./claerance`

#### Contribute
1. Fork it
2. Clone it: `git@github.com:Bermos/Claerance.git` or `https://github.com/Bermos/Claerance.git`
3. Create your feature branch: `git checkout -b my-new-feature`
4. Make changes and add them: `git add .`
5. Commit: `git commit -m 'Adds some feature'`
6. Push: `git push origin my-new-feature`
7. Pull request
