# Project's Purpose

This is a sample API created to test the process of developing and deploying APIs using golang and its tools. Created by [Hector Menezes](https://github.com/HectorMenezes).


# Local Setup


## Environment variables

Copy the following minimal information to a `.envfile`. The only missing variable will be the `SHORTENER_BASE_URL`.

```
POSTGRES_USER=test
POSTGRES_PASSWORD=test
POSTGRES_DB=url-shortener
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
```


## Running
To run locally, you'll need to run the database:

```shell
$ docker-compose up
```

Open another terminal and export the environment variables:

```shell
$ export $(cat .envfile | xargs)
```

Install requirements

```shell
$ go mod tidy
```

And then, run the API:

```shell
$ go run .
```
