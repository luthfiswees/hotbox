# Hotbox
Test Report Storage

## Dependencies

- Go 1.11 or newer
- MySQL
- S3 Compatible Object Storage

## Installation

- Install go 1.11 or newer. Refer [here](https://golang.org/doc/install).

- Install MySQL. Refer [here](https://dev.mysql.com/doc/mysql-getting-started/en/).

- Install Minio Server. Refer [here](https://docs.min.io/).

- Enable go modules

  ```sh
  export GO111MODULE=on
  ```

- Install dependencies

  ```sh
  go mod tidy
  ```

- Copy env

  ```sh
  cp en.sample .env
  ```

- Start service

  ```
  go run app/hotbox/main.go
  ```

## Endpoints

- **Store**

  Example Request (Postman cURL)

  ```sh
  curl -X POST \
    http://localhost:4001/store \
    -H 'Content-Type: ' \
    -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
    -F service_name=ayam \
    -F epoch_timestamp=123456789 \
    -F report=@/Users/luthfikurniaputra/Desktop/Placeholder/72x72.png \
    -F status=success
  ```

  Request Type : `form-data`

  Mandatory Body Values

  - `service_name` : Name of the service
  - `epoch_timestamp`: Timestamp in epoch values
  - `status` : Status of the test. 
  - `report`: Report File. Provide the directory path needed to find the file.

- **Get Report**

  Example Request

  ```sh
  curl -X GET \
    'http://localhost:4001/get-report?service_name=ayam&status=success'
  ```

  Mandatory Param Values:

  - `service_name`: Name of the service
  - `status`: Status of the report that you want to see

