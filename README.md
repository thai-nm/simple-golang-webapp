# Simple Webapp Golang
A simple Golang web application.

## API
| Endpoint | Type   | Usage             |
| -------- | ------ | ----------------- |
| `/tasks` | `GET`  | Get all tasks     |
| `/tasks` | `POST` | Create a new task |



## How to start
- Clone the repository and `cd` inside it:
  ``` bash
  git clone https://thainmuet@github.com/thainmuet/simple-golang-webapp.git
  cd simple-golang-webapp
  ```
- Build Golang code:
  ``` bash
  go build
  ```
- Start server:
  ``` bash
  ./simple-golang-webapp
  ```

## API call examples

- `POST` request:
  ``` bash
  curl --header "Content-Type: application/json" \
  --data '{"Name":"Project A", "Description":"A simple description"}' \
  -X POST http://localhost:8080/tasks
  ```

- `GET` request:
  ``` bash
  curl http://localhost:8080/tasks
  ```