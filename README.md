# Project's Purpose

This is a sample API created to test the process of developing and deploying APIs using golang and its tools. Created by [Hector Menezes](https://github.com/HectorMenezes).


# Local Setup


## Environment variables
### Database:
Use a `db.env` to declare the database environment variables:
```bash
POSTGRES_USER=test
POSTGRES_PASSWORD=test
POSTGRES_DB=url-shortener
```
### API
Use a `api.env` to declare the API environment variables. The databse information should be equal to the previous environment file:
```bash
POSTGRES_USER=test
POSTGRES_PASSWORD=test
POSTGRES_DB=url-shortener
POSTGRES_PORT=5432
POSTGRES_HOST=db
SHORTENER_BASE_URL=hm
```


## Running
To run locally, you'll only need to run with:

```bash
docker-compose up
```

## Testing
Once the API container is running, just execute
```bash
docker exec -it url-shortener-api go test ./...
```
