# Frozen Boars Hockey Club API
<p align="center">
  <img src="https://i.imgur.com/NXvdSJN.png">
</p>

## Installation
This API is in early development, and it's pretty much closed source so if you can see this then you've probably been given the link. If you're technically inclined and curious you can install docker and run `docker-compose run --service-ports fbhcapi` and be put into the shell of the container provided.

After you've been placed into the container, navigate to src/app/ and use `go get` to fetch dependencies.


## Usage
Not too many endpoints exist on this API as I'm just starting it, however this will change soon. If you have your own MySQL variables for db connection they will automatically be taken from OS environment if you use the following convention

- DB_USER=<your db username>
- DB_PASS=<your db password>
- DB_CONNSTRING=<your database url>
- DB_CONN_OPTIONALS=<any query params> eg: /query?param=blah *You don't need the slash*



## Features to be added / planned
- Registration for normal users and players that belong to the club
- Application process to apply to join the club (captains will have access)
- Drag and Drop system for captains to create lines easily, playes can log into the app and check lines

any other features you guys can think of, I'd be down to implement.


~ RSGG
