# golang-ip2location

For self hosted, ready to use RESTful API microservice for converting IP to location such a `City`, `Country`and `Region`. Free lite database version of `ip2location.com` is used. Please donwload your own database file from `ip2location.com`. You can also use `db-update.sh` script in the root directory for automatically refreshing database from `ip2location.com` with your own token. You can download databsae here [https://lite.ip2location.com/database-download](https://lite.ip2location.com/database-download)

Docker image: [https://hub.docker.com/r/vladimirok5959/golang-ip2location](https://hub.docker.com/r/vladimirok5959/golang-ip2location)

## Usage

```sh
Usage of ./ip2location:
  -access_log_file string
    Or ENV_ACCESS_LOG_FILE: Access log file
  -data_dir string
    Or ENV_DATA_DIR: Application data directory
  -db_update_time int
    Or ENV_DB_UPDATE_TIME: Delay in minutes between database reloading (default 60)
  -deployment string
    Or ENV_DEPLOYMENT: Deployment type (default "development")
  -error_log_file string
    Or ENV_ERROR_LOG_FILE: Error log file
  -host string
    Or ENV_HOST: Web server IP (default "127.0.0.1")
  -port string
    Or ENV_PORT: Web server port (default "8080")
  -web_url string
    Or ENV_WEB_URL: Web server home URL (default "http://localhost:8080/")
```

## API endpoint

Only one: http://localhost:8080/api/v1/ip2location/8.8.8.8

```txt
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 12 Oct 2022 20:45:48 GMT
Content-Length: 107
```

```json
{
    "City": "Mountain View",
    "CountryLong": "United States of America",
    "CountryShort": "US",
    "Region": "California"
}
```

## DB auto update

Right now, application is not designed for automatically refreshing database. Application just load pre-downloaded dabase file, load it once on startup and used it. So you must care about refreshing databse by yourself by using for example crontab and included `db-update.sh` script. Example for crontab file:

```sh
0    2    1    *    *    root    /var/ip2location/db-update.sh "your token" "/var/ip2location/data/IP2LOCATION-LITE-DB3.BIN" > /dev/null 2>&1
```

Script takes two parameters, first - your database token from `ip2location.com`, and second - database file. In this example script will automatically refresh database every month at 2 AM, once at month. It's enough for free lite database version. Script will not damage database file on fail

## Running docker container

```sh
docker run -d \
    --network host \
    --restart=always \
    --name my-container-name \
    -e ENV_DATA_DIR="/app/data" \
    -e ENV_DB_UPDATE_TIME="60" \
    -e ENV_DEPLOYMENT="deployment" \
    -e ENV_HOST="127.0.0.1" \
    -e ENV_PORT="8080" \
    -e ENV_WEB_URL="http://localhost:8080/" \
    -v /etc/timezone:/etc/timezone:ro \
    -v /var/ip2location/data:/app/data \
    vladimirok5959/golang-ip2location:latest
```
