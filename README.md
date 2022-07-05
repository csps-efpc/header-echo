# header-echo
Simple API in Go to echo a subset of HTTP headers back to the caller as a JSON message

## Building

```
go build
```

## Execution

```
header-echo -port [port] -pattern [pattern]
```
Where:
*  `[port]` is the TCP port on which the service should listen for HTTP requests (default 10001)
*  `[pattern]` is a regular expression that headers must match to be returned (default `^.*$`)

Calls should be made to `http://yourhost:[port]/echo`


-------------

# header-echo
API simple en Go pour renvoyer un sous-ensemble d'entêtes HTTP à l'appelant sous forme de message JSON.

## Construction

```
go build
```

## Exécution

```
header-echo -port [port] -pattern [pattern]
```
Où :
* `[port]` est le port TCP sur lequel le service écoutera les requêtes HTTP (par défaut 10001)
* `[pattern]` est une expression régulière à laquelle les en-têtes doivent correspondre pour être retournés (par défaut `^.*$`)

Les appels doivent être faits à `http://yourhost:[port]/echo`