FROM mariadb:10.5

# Override default "mysqld" cmd to set the some settings
CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci", "--max_allowed_packet=256MB"]
