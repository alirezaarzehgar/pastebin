package dotenv

const (
	KeyHttpAddress      = "HTTP_ADDRESS"
	KeyHttpPort         = "HTTP_PORT"
	KeyHttpIdleTimeout  = "HTTP_IDLE_TIMEOUT"
	KeyHttpReadTimeout  = "HTTP_READ_TIMEOUT"
	KeyHttpWriteTimeout = "HTTP_WRITE_TIMEOUT"

	KeyHandlerMaxFileSize = "HANDLER_MAX_FILE_SIZE"

	KeyLogLevel = "LOG_LEVEL"

	KeyObjectStorageMinIOEndpoint        = "OBJECT_STORAGE_MINIO_ENDPOINT"
	KeyObjectStorageMinIOAccessKeyID     = "OBJECT_STORAGE_MINIO_ACCESSKEYID"
	KeyObjectStorageMinIOSecretAccessKey = "OBJECT_STORAGE_MINIO_SECRETACCESSKEY"
	KeyObjectStorageMinIOUseSSL          = "OBJECT_STORAGE_MINIO_USESSL"

	KeyPersistenceCouchDBHostname       = "PERSISTENCE_COUCHDB_HOSTNAME"
	KeyPersistenceCouchDBPort           = "PERSISTENCE_COUCHDB_PORT"
	KeyPersistenceCouchDBUsername       = "PERSISTENCE_COUCHDB_USERNAME"
	KeyPersistenceCouchDBPassword       = "PERSISTENCE_COUCHDB_PASSWORD"
	KeyPersistenceCouchDBInsecure       = "PERSISTENCE_COUCHDB_INSECURE"
	KeyPersistenceCouchDBMetadataDBName = "PERSISTENCE_COUCHDB_METADATA_DB_NAME"
)
