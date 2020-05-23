# Gowiki

Simple static wiki engine written with golang. It uses markdown for articles and
converts to static html. Checkout [example](./example) folder.

## Run

In order to run example build wiki:
``` sh
$ go install github.com/i1i1/gowiki
```

Then `cd` into example directory and do:
``` sh
$ make run
```

It will launch apache httpd server in docker with all md pages inside [example/md](./example/md) directory
