package constant

const (
	EXPIRED_CACHE_CREDENTIAL = "expired_cache_credential"
	PROJECT_NAME             = "project_name"
	OWNER_NAME               = "Visionet"
	OWNER_EMAIL              = "visionetopsol@gmail.com"
	LAST_UPDATE              = "25-08-2022 17:41"

	SERVICE_ID      = "fs.api.template"
	SERVICE_NAME    = "template"
	SERVICE_VERSION = "1.0.3"

	CONFIG_PATH             = "config.json"
	CONFIG_PATH_DEVELOPMENT = "config-development.json"

	CONFIG_PATH_APP = "app"

	CONFIG_PATH_APP_DEBUG          = CONFIG_PATH_APP + ".debug"
	CONFIG_PATH_APP_HOST           = CONFIG_PATH_APP + ".host"
	CONFIG_PATH_APP_PORT           = CONFIG_PATH_APP + ".port"
	CONFIG_PATH_APP_USE_TLS        = CONFIG_PATH_APP + ".use_tls"
	CONFIG_PATH_APP_CERT_FILE_PATH = CONFIG_PATH_APP + ".cert_file_path"
	CONFIG_PATH_APP_KEY_FILE_PATH  = CONFIG_PATH_APP + ".key_file_path"
	CONFIG_PATH_APP_TIME_ZONE      = CONFIG_PATH_APP + ".timezone"

	CONFIG_PATH_CORS                    = CONFIG_PATH_APP + ".cors"
	CONFIG_PATH_ALLOW_ORIGIN            = CONFIG_PATH_CORS + ".allow_origin"
	CONFIG_PATH_ALLOW_METHOD            = CONFIG_PATH_CORS + ".allow_method"
	CONFIG_PATH_ALLOW_HEADER            = CONFIG_PATH_CORS + ".allow_header"
	CONFIG_PATH_EXPOSE_HEADER           = CONFIG_PATH_CORS + ".expose_header"
	CONFIG_PATH_ALLOW_CREDENTIAL        = CONFIG_PATH_CORS + ".allow_credential"
	CONFIG_PATH_ALLOW_WILDCARD          = CONFIG_PATH_CORS + ".allow_wildcard"
	CONFIG_PATH_ALLOW_BROWSER_EXTENSION = CONFIG_PATH_CORS + ".allow_browser_extension"
	CONFIG_PATH_ALLOW_WEB_SOCKET        = CONFIG_PATH_CORS + ".allow_web_socket"
	CONFIG_PATH_ALLOW_FILE              = CONFIG_PATH_CORS + ".allow_file"
	CONFIG_PATH_MAX_AGE                 = CONFIG_PATH_CORS + ".max_age"

	CONFIG_PATH_DOMAIN    = "domain"
	CONFIG_PATH_IP_SERVER = "ip_server"
	CONFIG_PATH_DB        = "database"

	CONFIG_PATH_DB_HOST     = CONFIG_PATH_DB + ".host"
	CONFIG_PATH_DB_PORT     = CONFIG_PATH_DB + ".port"
	CONFIG_PATH_DB_USERNAME = CONFIG_PATH_DB + ".username"
	CONFIG_PATH_DB_PASSWORD = CONFIG_PATH_DB + ".password"
	CONFIG_PATH_DB_DATABASE = CONFIG_PATH_DB + ".database"
	CONFIG_PATH_DB_DIALECT  = CONFIG_PATH_DB + ".dialect"
	CONFIG_PATH_DB_SSL_MODE = CONFIG_PATH_DB + ".sslmode"

	CONFIG_PATH_REDIS = "redis"

	CONFIG_PATH_REDIS_HOST     = CONFIG_PATH_REDIS + ".host"
	CONFIG_PATH_REDIS_PASSWORD = CONFIG_PATH_REDIS + ".password"
	CONFIG_PATH_REDIS_DB       = CONFIG_PATH_REDIS + ".db"

	CONFIG_PATH_SMTP = "smtp"

	CONFIG_PATH_SMTP_HOST     = CONFIG_PATH_SMTP + ".host"
	CONFIG_PATH_SMTP_PORT     = CONFIG_PATH_SMTP + ".port"
	CONFIG_PATH_SMTP_EMAIL    = CONFIG_PATH_SMTP + ".email"
	CONFIG_PATH_SMTP_PASSWORD = CONFIG_PATH_SMTP + ".password"
	CONFIG_PATH_SMTP_NAME     = CONFIG_PATH_SMTP + ".name"

	CONFIG_MQTT_BROKER         = "mqtt_broker"
	CONFIG_PATH_ATOME          = "atome"
	CONFIG_PATH_ATOME_URL      = CONFIG_PATH_ATOME + ".url"
	CONFIG_PATH_ATOME_USERNAME = CONFIG_PATH_ATOME + ".username"
	CONFIG_PATH_ATOME_PASSWORD = CONFIG_PATH_ATOME + ".password"

	DEFAULT_APP_HOST  = "localhost"
	DEFAULT_APP_PORT  = "8080"
	DEFAULT_APP_DEBUG = false

	DEFAULT_DB_HOST     = "localhost"
	DEFAULT_DB_PORT     = "3306"
	DEFAULT_DB_USERNAME = "root"
	DEFAULT_DB_PASSWORD = ""
	DEFAULT_DB_DATABASE = "template"
	DEFAULT_DB_DIALECT  = DIALECT_MYSQL
	DEFAULT_DB_SSL_MODE = "disable"
	DEFAULT_TIME_ZONE   = "Asia/Jakarta"
	DEFAULT_PATH_SQLITE = SERVICE_NAME + ".db"

	DEFAULT_REDIS_HOST     = "localhost:6379"
	DEFAULT_REDIS_PASSWORD = ""
	DEFAULT_REDIS_DB       = 0

	DEFAULT_SMTP_HOST     = "smtp.gmail.com"
	DEFAULT_SMTP_PORT     = 587
	DEFAULT_SMTP_EMAIL    = "sdh.notification@gmail.com"
	DEFAULT_SMTP_PASSWORD = "Visionet*1!"
	DEFAULT_SMTP_NAME     = "San Diego Hills Notification"

	DEFAULT_ONE_SIGNAL_APP_ID       = "c90ea719-f8cb-4eb9-b946-bb6516927b94"
	DEFAULT_ONE_SIGNAL_REST_API_KEY = "MGM3MzQwMmUtYTliOC00NmU3LWI2NGMtMTNlNmM0ZWFiNjU1"

	DEFAULT_NEW_RELIC_APP_NAME = "sdh"
	DEFAULT_NEW_RELIC_LICENSE  = "0adb99dfc961708a1218481e50a266b49f9bNRAL"

	DEFAULT_MINIO_HOST           = "localhost:9000"
	DEFAULT_MINIO_LOCATION       = "us-east-1"
	DEFAULT_MINIO_ACCESS_KEY     = "4JFFD40NAS1768JPWVVV"
	DEFAULT_MINIO_SECRET_KEY     = "JBaIB6bwL2CDWD1hoDT6t7q+zKVQv+TBH8Cx+jhR"
	DEFAULT_MINIO_SSL            = false
	DEFAULT_MINIO_REPLACE_DOMAIN = ""

	MAX_LIMIT_QUERY     = 100
	DEFAULT_LIMIT_QUERY = 10

	DIALECT_MYSQL      = "mysql"
	DIALECT_POSTGRESQL = "postgres"
	DIALECT_SQL_SERVER = "mssql"
	DIALECT_SQLITE     = "sqlite"

	SWAGGER_HOST = "swagger_host"
)
