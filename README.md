# btc-alert-api
An API to alert users when BTC reaches their expected price

## SQLite3 Table Create
```
CREATE TABLE "alertUsers" (
	"email"	INTEGER NOT NULL UNIQUE,
	"priceToCheck"	INTEGER NOT NULL,
	"status"	TEXT NOT NULL DEFAULT 'active',
	PRIMARY KEY("email")
);
```