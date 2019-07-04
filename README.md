# Requirements

## Development
- `go get github.com/codegangsta/gin`
- https://github.com/rvoicilas/inotify-tools/wiki

# RUN for Development 

Run `$ while inotifywait -e modify pkg/dbuser/*; do touch cmd/main.go; done` in a terminal

In other terminal run `$GOPATH/bin/gin --path cmd --port 3000`