#Nats go cli client

## Usage ##

### Subscribing ###

queueGroup is optional
```
    ./natsCli -host http://localhost:4222 -type subscribe -topic MyTopic -queueGroup MyQueue
```

### Publishing ###
```
    ./natsCli -host http://localhost:4222 -type publish -topic MyTopic -data MyData
```


## Building ##
```
    go build -o natsCli natsCli.go
```