# Purpose
Program written in GO which spins up a web server, it prints the hostname and
returns a Code500 every 5 hits

#Usage
To get source then compile it:

    $ go get github.com/richardpct/go-hostname

To launch the webserver:

    $ ~/go/bin/go-hostname

To see the output:

    $ curl localhost:8080
