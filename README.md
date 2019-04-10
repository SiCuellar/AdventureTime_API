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
go mod download``
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

#### `POST /api/v1/login?username=<USERNAME>`
This endpoint will return the user object associated with <USERNAME>.

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

#### `POST /api/v1/quest?user_id=<USER_ID>&lat=<LATITUDE>&long=<LONGITUDE>`
This endpoint will create a new quest with provided lat and long, or return the currently active one if it exists for a given user.

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "Location1": "4a8b585ff964a520360c20e3|1822 Blake St (btwn 19th St & 18th St), Denver, CO 80202, United States",
    "Location2": "4e10f909483bee47ff2e50c0|Denver, CO 80204, United States",
    "Location3": "53220885498e6416b2ed973a|1433 17th St (Blake), Denver, CO 80202, United States",
    "Status": 0,
    "User": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "username": "",
        "current_hp": 0,
        "current_xp": 0,
        "items": null
    },
    "UserID": 1
}
```

If a lat/long is not provided, the following json will be returned:

```json
{
    "Error": "You must provide a lat and long"
}
```

#### `POST /api/v1/checkin?user_id=<USER_ID>&lat=<LATITUDE>&long=<LONGITUDE>`
This endpoint will match the lat/long provided to a set of likely foursqaure location IDS. A successful match will return:

```json
{
    "success": "Lat/Long matches current goal location."
}
```
An unsuccessful match will return:
```json
{
    "error": "Lat/Long does not match current goal location."
}
```
#### `POST /api/v1/encounter?success=<TRUE/FALSE>&user_id=<USER_ID>`
This endpoint will add xp and update location if the encounter is successful and will return:

```json
{
    "success": "Succesful Encounter."
}
```
An unsuccessful match will reset XP,HP and quest to initial starting values and will return:
```json
{
    "error": "Encounter Failed."
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
