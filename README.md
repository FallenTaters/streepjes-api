# Streepjes App Backend

Streepjes app is used for keeping track of products ordered by members of De Parabool or Lacrosse Groningen Gladiators. All data is stored in a file called `streepjes.db` which should be backed up regularly.

## Usage
### Dependencies
required
- go
- enumer (https://github.com/alvaroloes/enumer)

optional
- make
- air (https://github.com/cosmtrek/air)

### Get Started - Production
- configure environment (see below)
- generate enumer files
    - `make generate`
- make a production build
    - `make build`
- create streepjes.db **if it doesn't exist**
    - `make resetdb`
- run the binary
    - `./bin/streepjes`

### Get Started - Development
- configure environment (see below)
- generate enumer files
    - `make generate`
- create streepjes.db
    - `make resetdb`
- `make run`

### Env settings
- PORT (default 81, use 8081 for development)
- DISABLE_SECURE (default false, use true for development)

### Calls
- An `openapi.yml` will be included soon (hopefully).

## Further Information
The database file is managed using bbolt key-value sture (https://github.com/etcd-io/bbolt), specifically with my personal wrapper for it, bbucket (https://github.com/PotatoesFall/bbucket). All data is encoded as JSON.

All routes are defined in `./cmd/streepjes/main.go`, and handled in `./cmd/streepjes/route.go`. Further logic is handed down by domain in `./domain/...`. Public functions should be placed in an eponymous file, i.e. `./domain/members/members.go`, and all interactions with the database should be handled in a file called repo.go, i.e. `./domain/members/repo.go`

`./cmd/maketestdb` is used for generating a new streepjes.examples.db after changes in development. Note that when changing the types written to `streepjes.db`, they must be backwards compatible. A migration module may be added later to support this.

Orders are linked to a member and a user (bartender), and contain a JSON string of the products and their prices. This way, if products or their prices are mutated, existing orders are not affected. The order domain currently imports the other packages, but the other domains may not import the order package, to avoid cyclical import. The main package should handle cross-domain logic.