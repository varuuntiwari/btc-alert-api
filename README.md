# btc-alert-api
An API to alert users when BTC reaches their expected price

## Endpoints
### /alerts/create
Takes query parameters "email" and "priceCheck" to create an alert
### /alerts/delete
Takes a parameter "email" to delete the alert the email ID is associated with

## Usage
Run the file `btc-alert.go` and test API using the endpoints given above
