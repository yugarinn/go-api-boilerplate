#!/bin/bash

TEST_DATABASE="$MYSQL_DATABASE"_test

echo "CREATE DATABASE IF NOT EXISTS \`$TEST_DATABASE\` ;" | "${mysql[@]}"

# If a user was configured, grant him access to the test database
if [ "$MYSQL_USER" -a "$MYSQL_PASSWORD" ]; then
    echo "GRANT ALL ON \`$TEST_DATABASE\`.* TO '$MYSQL_USER'@'%' ;" | "${mysql[@]}"
fi

echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"
