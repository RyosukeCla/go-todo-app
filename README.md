# Go todo app
simple todo app for learning go.
no design, just it.

## stack
- react 16.5.2
- go 1.11.1
- scribble: a tiny golang json db

## start server
### mac: prebuild binary
```bash
$ ./go-todo-app
```

### install requirements
```bash
$ sh ./scripts/install.sh
```

### dev
```bash
$ godo dev --watch
$ cd ./client; npm run dev
# go to localhost:3000
```

### build
```bash
$ sh ./scripts/build.sh
$ ./go-todo-app
```
