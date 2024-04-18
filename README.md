# IPFS File API
A backend server provides an API for uploading and downloading files to IPFS, and stores the data in a PostgreSQL database.

## Overview

- Language: Go v1.21.4
- Web FrameWork: Gin v1.9.1
- PostgreSQL v15.1
- ipfs v0.27.0

## Run

### Update Modules
```
go get -u && go mod tidy -v
```


### Run
```
go run cmd/main.go
```

## ENV
copy .env.default and rename as .env
```
API_PORT=8080
POSTGRES_HOST=
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_PASSWORD=password
POSTGRES_DB=
IPFS_HOST=ipfs:5001
```

## API

- GET /api/file-cid/{id}: 取得檔案
- POST /api/file-cid/{name}: 上傳檔案
