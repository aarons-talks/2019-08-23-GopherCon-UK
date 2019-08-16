# Building An App with Go Modules and Athens :tada:

Hey Gophers! We're gonna build an awesome webapp with [Gin](https://github.com/gin-gonic/gin). Well, it's pretty basic but it shows cat and dog pictures so it's still pretty awesome :grinning:.

The cool part though? Instead of pulling all my webapp's dependencies directly from version control systems like GitHub (which we've always been doing in the past), we're gonna build it using module proxies and [Athens](https://docs.gomods.io).

In fact, we're actually gonna do the build in two different ways! We'll talk about each as we go. I developed and tested this demo on a Zsh shell on MacOS.

>I gave this demo at [GopherCon UK 2019](https://www.gophercon.co.uk/) in the session titled " Improving Dependencies for Everyone with Modules, Proxies and the Athens Project".

![athens banner](./athens-banner.png)

# About the Web App

The web application we're going to build is fairly simple, but not trivial. I built a little server using [gin](https://github.com/gin-gonic/gin) as the framework, and it shows some HTML pages with cat pictures on them.

The cool thing is that the whole codebase is modules-aware. That means it has a `go.mod` file that keeps track of all my code's dependencies. The built-in [Go Modules](https://github.com/golang/go/wiki/Modules) system reads that file to look up the dependencies it needs to download before it builds my server.

# Run The Demo

Below is how how to do the demo yourself. The instructions are for Linux/Mac OS X systems.

## First Way: Build With Athens and an Upstream VCS

Athens starts up initially with nothing in its storage. When you run `go get`, it downloads modules from an "upstream". In this demo, we're configuring it to fetch code directly from the VCS, and then store it forever. You can also configure it to download from module mirrors like proxy.golang.org or gocenter.io.

### Run The Server!

We try hard to make it easy to run your own Athens. See [here](https://docs.gomods.io/install) for instructions for running the server a few different ways. Today, we're going to use [Docker](https://www.docker.com/) to run ours.

First, run this to start Athens up:

```console
$ docker run -p 3000:3000 -e GO_ENV=development -e ATHENS_GO_GET_WORKERS=5 gomods/athens:v0.5.0
```

And then to set your `GOPROXY` environment variable to tell modules to use the local server:

```console
$ export GOPROXY=http://localhost:3000
```

Also, the Go tool keeps a read-only on-disk cache of every module version you've downloaded for any build. To make it read-only, it stores each file in the cache with `-r--r--r--` permissions. Since that's the case, you need to use `sudo` to clear the cache.

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
```

And then build and run the server!

```console
$ go run .
```

## Second Way: Use Your Athens While Offline :scream:

Like I mentioned in the last demo, Athens stores the dependencies you use _forever_ in its own storage. That means that you can build your code without access to the internet. Let's do that here!

In the previous step, Athens was using in-memory storage, which is strictly for local development / demonstration purposes.

In this demo, we're going to run Athens using its disk storage driver, and pre-load its module database - located in this repository at `athens-archive/` with all the dependencies our app needs. This way, Athens will be able to serve all the dependencies that `go run` requests, and won't ever need to fetch code from the public internet.

First, run Athens with the disk driver, and mount the module database into the Docker container:

```console
$ export ATHENS_ARCHIVE="$PWD/athens-archive"
$ docker run -p 3000:3000 -e GO_ENV=development -e ATHENS_GO_GET_WORKERS=5 -e ATHENS_STORAGE_TYPE=disk -e ATHENS_DISK_STORAGE_ROOT=/athens -v $ATHENS_ARCHIVE:/athens gomods/athens:v0.5.0
```

>Set `ATHENS_ARCHIVE` to wherever your archive lives. If, for example, it is on a USB key, set it to `/Volumes/MyKey` (this is where it would most likely live on a Mac)

Next, clear out your cache again:

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
```

And then, **shut down your internet connection** :see_no_evil:.

And finally, run the app again!

```console
$ go run .
```

>Notice how much faster the build is this time, compared the the previous demo. That's because Athens isn't downloading anything - it's just streaming data directly from disk to the client!

### If you want to create your own disk archive

Athens treats storage like a cache, except without the evictions or TTLs. That means you can mount an empty volume, connect to the internet, and execute `go run` (or similar) against that Athens.

On every "miss", Athens will _synchronously_ download the module and store it. After your `go run` succeeds, Athens guarantees that it will have stored all your app's dependencies on disk (the same guarantee applies to all other storage drivers too!).

These commands will set up Athens with disk storage and an empty volume:

```console
$ mkdir $ATHENS_ARCHIVE
$ docker run -p 3000:3000 -e GO_ENV=development -e ATHENS_GO_GET_WORKERS=5 -e ATHENS_STORAGE_TYPE=disk -e ATHENS_DISK_STORAGE_ROOT=/athens -v $ATHENS_ARCHIVE:/athens gomods/athens:v0.5.0
```

After you run them, set `GOPROXY` and do a `go run` like we did in the previous demos, and the disk archive will be at `$ATHENS_ARCHIVE`.

## Finally

Thanks for following along! Now you're both a Gopher _and_ an Athenian :green_heart:.

If you want to learn more, check out [docs.gomods.io](https://docs.gomods.io)! We'd also love for you to get involved - here are some ways to do so:

- Come star [our repo](https://github.com/gomods/athens)
- Come say hello on the `#athens` channel in the [Gophers Slack](https://invite.slack.golangbridge.org/)
  - This is a great place to come ask for help getting started and ask questions too :smile:
- And of course, file [bug reports](https://github.com/gomods/athens/issues/new/choose) or [feature requests](https://github.com/gomods/athens/issues/new/choose) and [contribute code](https://docs.gomods.io/contributing/new/development/)!

# Keep on rockin', Gophers!

![athens gopher](./athens-gopher.png)
