#Nats go cli client

## Usage ##

You'll need to run NATs somehow, for ease of use, you can use Docker.
```
    docker run -dP nats -user=foo -pass=bar
```

### Subscribing ###

queueGroup is optional
```
    ./natsCli -host http://localhost:4222 -type subscribe -topic MyTopic -queueGroup MyQueue
    ./natsCli -host http://foo:bar@localhost:4222 -type subscribe -topic MyTopic -queueGroup MyQueue
```

### Publishing ###
```
    ./natsCli -host http://localhost:4222 -type publish -topic MyTopic -data MyData
    ./natsCli -host http://foo:bar@localhost:4222 -type publish -topic MyTopic -data MyData
```


## Building ##
```
    go get github.com/nats-io/nats
    go build -o natsCli natsCli.go
```
