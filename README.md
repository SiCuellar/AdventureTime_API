# AdventureTime API

AdventureTime

Deployment URL: https://adventure-time-m4cap.herokuapp.com/

## Installation

Grab the source code and `cd` into the directory:

```bash
git clone https://github.com/SiCuellar/AdventureTime_API.git
cd AdventureTime_API
```

Then download all the dependencies and compile an executable:

```bash
go mod download
go build
```

## Usage

### Environment Variables
You must have the following environment variables set:

- `FOUR_ID` - Fourquare API Client ID
- `FOUR_SECRET` - Foursquare API Client Secret
- `PORT` - Port for the webserver to listen on


### Running the Server
To start the API server in development:
```bash
go run main.go
```

To start the API server in production:
```bash
./AdventureTime_API
```

### API Endpoints

#### `POST /api/v1/login?user_id=<USER_ID>`
This endpoint will return the user object associated with <USER_ID>.

Example Response:
```json
{
  "Value": {
      "ID": 1,
      "CreatedAt": "2019-04-04T11:44:19.311179-06:00",
      "UpdatedAt": "2019-04-04T11:44:19.311179-06:00",
      "DeletedAt": null,
      "username": "cebarks",
      "current_hp": 0,
      "current_xp": 0,
      "items": [
          {
              "ID": 1,
              "CreatedAt": "2019-04-04T13:09:17.894602-06:00",
              "UpdatedAt": "2019-04-04T13:09:17.894602-06:00",
              "DeletedAt": null,
              "name": "Test Item",
              "attack": 10,
              "defense": 10
          }
      ]
  },
  "Error": null,
  "RowsAffected": 1
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
