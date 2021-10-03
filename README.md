# Food delivery back-end
(for studying golang)

## Quick start

1. Install docker, golang.
2. Create a new .env file in the project folder following this example:
    ```
        MYSQL_VERSION=8.0
        DB_NAME=DB
        DB_USER=root
        DB_PASSWORD=root
        DB_ROOT_PASSWORD=root
        DB_PORT=3306
        DB_URI=root:root@/DB?parseTime=true
        VOLUME_PATH=/Users/<username>/Workplace/mysql
        GOOGLE_APPLICATION_CREDENTIALS=<FILE_PATH>
        GOOGLE_CLOUD_STORAGE_BUCKET_NAME=<BUCKET_NAME>
        GOOGLE_CDN=<URL>
        SECRET_KEY=<KEY>
        ACCESS_TOKEN_EXPIRE_TIME=604800
        REFRESH_TOKEN_EXPIRE_TIME_KEY=1209600
    ```
3. Run ```docker-compose up``` to start a mysql instance
4. ```go run main.go```
