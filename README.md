# Frozen Boars Hockey Club API
<p align="center">
  <img src="https://i.imgur.com/NXvdSJN.png">
</p>

## Installation
This API is in early development, and it's pretty much closed source so if you can see this then you've probably been given the link. If you're technically inclined and curious you can install docker and run `docker-compose up` and the api should be listening on port 8080.


## Usage
Not too many endpoints exist on this API as I'm just starting it, however this will change soon. If you have your own MySQL variables for db connection they will automatically be taken from OS environment if you use the following convention

| Variable | Variable Entry | Example |
|----------|----------------|---------|
| DB_USER  | your username  | DB_USER=username |
| DB_PASS  | your password  | DB_PASS=secret |
| DB_CONNSTRING | your database url| DB_CONNSTRING=host.service.whatever|
| DB_CONN_OPTIONALS | any query params | /query?param=blah *You don't need the slash* |

## Features to be added / planned
- Registration for normal users and players that belong to the club
- Application process to apply to join the club (captains will have access)
- Drag and Drop system for captains to create lines easily, playes can log into the app and check lines

any other features you guys can think of, I'd be down to implement.


~ RSGG
