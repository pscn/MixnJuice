# Running

`docker-compose up --build`

If you added modules, do `dep ensure` first.

Use  `docker-compose rm` to get rid of old data volumes.

Compiled in one command line it looks like this:

```sh
dep ensure && yes | docker-compose rm && docker-compose up --build
```

# Notes

Uses dep for dependency managment.

Install with: `go get -u github.com/golang/dep/cmd/dep`

Then run `dep ensure` in this directory to populate the `vendor` directory.

# go.mongodb.org/mongo-driver

I played with this one first.  It's too low level to get any work done.