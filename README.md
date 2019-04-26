# studentjobengine

a location based student jobs api

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make build
$ ./bin/studentjobengine
```

### Testing

``make test``

### Tips

* ``curl -H "Content-Type: application/json" --data @sample.json 127.0.0.1:8080/jobs/add``
* ``curl 127.0.0.1:8080/jobs | python -m json.tool``

### Login

* ``curl -H "Content-Type: application/json" --data @user_login.json 127.0.0.1:8080/users/login |python -m json.tool``
* ``export TOKEN=<JWT>``
* ``curl -H "Authorization: Bearer $TOKEN" <url>``
