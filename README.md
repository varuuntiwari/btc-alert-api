# btc-alert-api
An API to alert users when BTC reaches their expected price.

# Endpoints
## /alerts/create
Takes query parameters "email" and "priceCheck" to create an alert.
## /alerts/delete
Takes a parameter "email" to delete the alert the email ID is associated with.

# Usage
1. Set an env variable `MONGO_URI` with the URL to the database.
2. Make sure the port 8080 is free to use.
3. Run the command `go mod download && go run btc-alert.go`.

# Docker Compose
Run the command `docker-compose up -d` at the root of the repo to run the app in detached mode. This will create two containers, one from `mongo:latest` and the other from `golang:latest` image. The port 8080 and 27017 is accessible on the host machine.

# Tasks left
- [x] Function to check BTC price at intervals.
- [ ] Function to search for matching prices in the collection with the `status` as 'active'.
- [ ] Make the function run asynchronously and call the alert accordingly.