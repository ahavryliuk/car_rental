#!/bin/bash

chmod +x /cr-bin/rental_api
chmod +x /cr-bin/bookings_logger

until nc -z -v -w30 db 3306
do
  echo "Waiting for database connection..."
  sleep 3
done
echo "MySQL is up"

until nc -z -v -w30 redis 6379
do
  echo "Waiting for Redis..."
  sleep 3
done
echo "Redis is up"

until nc -z -v -w30 server 80
do
  echo "Waiting for Nginx..."
  sleep 3
done
echo "Nginx is up"

cd /cr-bin; (trap 'kill 0' INT; ./rental_api & ./bookings_logger)