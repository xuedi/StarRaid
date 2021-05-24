#!/bin/bash
SQL_USER=$(sed -n 's/.*user *= *\([^ ]*.*\)/\1/p' < config.ini)
SQL_PASS=$(sed -n 's/.*pass *= *\([^ ]*.*\)/\1/p' < config.ini)
SQL_NAME=$(sed -n 's/.*name *= *\([^ ]*.*\)/\1/p' < config.ini)
SQL_HOST=$(sed -n 's/.*host *= *\([^ ]*.*\)/\1/p' < config.ini)
SQL_PORT=$(sed -n 's/.*port *= *\([^ ]*.*\)/\1/p' < config.ini)
SQL="mysql --user=$SQL_USER --password=$SQL_PASS --database=$SQL_NAME --host=$SQL_HOST --port=$SQL_PORT"

$SQL starraid < seed.sql

