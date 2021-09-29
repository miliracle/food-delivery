# Food delivery back-end
(for studying golang)

## Quick start

1. Install docker, golang.
2. Create a new .env file in the project folder following this example:
    ```
        MYSQL_VERSION=5.7
        DB_NAME=DB
        DB_USER=root
        DB_PASSWORD=root
        DB_ROOT_PASSWORD=root
        DB_PORT=3306
        DB_URI=root:root@/DB
        VOLUME_PATH=/Users/<username>/Workspace/mysql
        GOOGLE_APPLICATION_CREDENTIALS=credential/hobby-323104-ce47c45d29de.json
        GOOGLE_CLOUD_STORAGE_BUCKET_NAME=buket-miliracle
        GOOGLE_CDN=abc.cdn.com
    ```
3. Run ```docker-compose up``` to start a mysql instance
4. ```go run main.go```
