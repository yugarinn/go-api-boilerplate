FROM mysql:8.0

# Configure a second testing database
COPY ./setup-test-db.sh /docker-entrypoint-initdb.d/

# Override default "mysqld" cmd to set the some settings
CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci", "--max_allowed_packet=256MB", "--default-authentication-plugin=mysql_native_password"]
