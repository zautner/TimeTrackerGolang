# TimeTrackerGolang

## Time tracker for a random NFC clock with optional functionality

### Runs on port 8080 of the localhost

### Endpoints:

#### Tracker functionality

* POST  "/track/:name"
    - track user activity
* PUT   "/track/:name"
    - reset user state
* GET   "/track/days/:name/:back"
    - get user records for the last :back days
* GET   "/track/months/:name/:back"
    - get user records for the last :back months

#### User store

* GET       "/user/:name"
    - Get user data for the currenrt month
* POST      "/user/:name"
    - Explicit add user
* PUT       "/user/:name/:state"
    - Explicit change user state
* DELETE    "/user/:name"
    - Delete user
* GET       "/user"
    - List users

#### Command line testing

` go test -v ./services` from the current folder