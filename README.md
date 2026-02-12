# Pastebin

A simple pastebin service written in Go.
It allows uploading text content along with optional files, storing file data in object storage and metadata in a persistence layer. The service supports configurable expiry for each paste.

## Features

* Upload text content and optional files in a single request
* Store file data in MinIO (S3-compatible object storage)
* Store metadata in CouchDB
* Cache support using Redis
* Configurable via `.env`
* HTTP API

## Requirements

* Go 1.21+
* Docker and Docker Compose (recommended)
* CouchDB
* MinIO
* Redis

You can either run dependencies manually or use Docker Compose.

## Configuration

Configuration is loaded from a `.env` file.

Create a `.env` file in the project root (you can copy from `.env.example` if available):

```bash
cp .env.example .env
```

Then adjust the values based on your environment.

### HTTP Configuration

* `HTTP_ADDRESS` – HTTP bind address (e.g., `0.0.0.0`)
* `HTTP_PORT` – HTTP port (e.g., `8080`)
* `HTTP_IDLE_TIMEOUT` – idle timeout in seconds
* `HTTP_READ_TIMEOUT` – read timeout in seconds
* `HTTP_WRITE_TIMEOUT` – write timeout in seconds

### Handler

* `HANDLER_MAX_FILE_SIZE` – maximum upload file size in bytes

### Logging

* `LOG_LEVEL` – logging level (e.g., `debug`, `info`)

### MinIO (Object Storage)

* `OBJECT_STORAGE_MINIO_ENDPOINT` – MinIO endpoint (e.g., `127.0.0.1:9000`)
* `OBJECT_STORAGE_MINIO_ACCESSKEYID` – access key
* `OBJECT_STORAGE_MINIO_SECRETACCESSKEY` – secret key
* `OBJECT_STORAGE_MINIO_USESSL` – use SSL (`true` or `false`)
* `MINIO_DEFAULT_BUCKETS` – comma-separated bucket names

### CouchDB (Persistence)

* `PERSISTENCE_COUCHDB_HOSTNAME` – CouchDB host
* `PERSISTENCE_COUCHDB_PORT` – CouchDB port
* `PERSISTENCE_COUCHDB_USERNAME` – username
* `PERSISTENCE_COUCHDB_PASSWORD` – password
* `PERSISTENCE_COUCHDB_INSECURE` – skip TLS verification
* `PERSISTENCE_COUCHDB_METADATA_DB_NAME` – metadata database name

### Redis (Cache)

* `KEY_CACHE_REDIS_ADDRESS` – Redis address (e.g., `127.0.0.1:6379`)
* `KEY_CACHE_REDIS_PASSWORD` – Redis password
* `KEY_CACHE_REDIS_DB` – Redis DB index

## Running with Docker Compose

Start all dependencies:

```bash
docker compose up -d
```

Ensure your `.env` configuration matches the service addresses defined in `docker-compose.yaml`.

## Running the Application

Download dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run ./cmd/pastebin
```

The server will start on the configured address and port.

## API

### Create Paste

`POST /paste`

Form fields:

* `content` – text content
* `file` – file upload (optional, can be sent multiple times if supported)
* `expiry` – duration string (e.g., `60s`, `10m`, `1h`)

Example:

```bash
curl -X POST http://localhost:8080/paste \
  -F "file=@/etc/passwd" \
  -F "content=yooooooooooo yooooooooooooo yooooooooo" \
  -F "expiry=60s"
```

Response:

```json
{
  "paste_id": "ca596523-d586-450b-985f-828fd2140c24"
}
```

### Get Paste Content

`GET /paste/{id}`

Returns the paste content and metadata of attached files.

Example:

```bash
curl -X GET http://localhost:8080/paste/ca596523-d586-450b-985f-828fd2140c24
```

Response:

```json
{
  "content": "yooooooooooo yooooooooooooo yooooooooo",
  "metadatas": [
    {
      "filename": "passwd",
      "content-type": "application/octet-stream",
      "size": 4342,
      "checksum": "f9aa5a4cbb1cac1850f94e1b2f2e70e4"
    }
  ]
}
```

### Download Paste File

`GET /paste/{id}/{filename}`

Downloads a specific file attached to the paste.

Example:

```bash
curl -O http://localhost:8080/paste/ca596523-d586-450b-985f-828fd2140c24/passwd
```

If the paste has expired, it will no longer be accessible.

## Testing

After running the server:

1. Send a `POST /paste` request to create a new paste.
2. Copy the returned `paste_id`.
3. Send a `GET /paste/{id}` request to retrieve its content.
4. Optionally download a file using `GET /paste/{id}/{filename}`.

You can use `curl`, Postman, or any HTTP client to test the API.

