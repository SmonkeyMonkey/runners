# _Application for marathon runners_
Features:
- Get the list of all runners (GET /runner)
- Get a single runner with all results (GET /runner/:id)
- Get the top 10 runners for the selected country (GET  /runner?country=value)
- Get the top 10 runners for the selected season (year)  (GET /runner?year=value OR by country&year /runner?country=value&year=value)
- Create a new runner (POST /runner)
- Update existing runner (PUT /runner)
- Add (create) a new result (POST /result)
- Remove (delete) existing results (DELETE /result/:id)
- Delete runner (DELETE /runner/:id)
- Login (POST /login)
- Logout (POST /logout)
- Metrics ( localhost:9000/metrics )

all routes uses port 8080 excepts metrics
For visualisation metrics you can use Grafana
###  Installation
Clone repository:
```bash
git clone github.com/smonkeymonkey/runners.git
```
### Usage
Run locally via docker-compose:
```bash
docker-compose up -d
```
### Login
Example Postman:
Basic auth
Login: admin 
Password: admin
It returns a token to be inserted into the `token` header.