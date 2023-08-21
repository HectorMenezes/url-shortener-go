#!/bin/sh

until psql postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -c "\q"; do
  sleep 5
done
swag init
/url-shortener
