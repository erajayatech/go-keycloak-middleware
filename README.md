# Purpose
Specific JWT Validation Middleware for Keycloak Authorization Permission Scope.

# Support
* Go Application with Gin Web Framework.
* JWT Validation with Json Web Key(JWK).

# Dependency
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [JWT by cristalhq](https://github.com/cristalhq/jwt)
* [GoDotEnv by joho](https://github.com/joho/godotenv)

# Install
Go Version 1.15+
```
go get github.com/cakobob/go-keycloak-middleware
```

# Setup Environment
* KEYCLOAK_JWT_ISS
```.env
example in .env file: 
KEYCLOAK_JWT_ISS=http://localhost:9999/auth/realms/dashboard
```
* KEYCLOAK_JWT_JWK_ENDPOINT
```.env
example in .env file: 
KEYCLOAK_JWT_JWK_ENDPOINT=http://localhost:9999/auth/realms/dashboard/protocol/openid-connect/certs
```

# Example: Keycloak Token
```
eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJtU0czMFVkV3FfZU5XLU1PdEdSWWxrVkx1Z3RjbjA5NTJfU3BJc04xX0RVIn0.eyJleHAiOjE2NDA4Mzk0MTQsImlhdCI6MTY0MDgzOTExNCwianRpIjoiYzFjMjcwMTItMzI1Zi00ZjVhLTgzZWQtNTI5MGU1MjExZDBiIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5OTk5L2F1dGgvcmVhbG1zL2Rhc2hib2FyZCIsImF1ZCI6ImRhc2hib2FyZC1hcGkiLCJzdWIiOiIxZGZkYjljMi0xMmU2LTRlNWYtYWRmOC02OWE0Y2UxZjI0ZGIiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJkYXNoYm9hcmQtYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjJkZDEwMTA1LTNhNzctNDUyNS1iMmMxLTVmNmNmYjA3MTU2NSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib3JkZXJfdmlld2VyIl19LCJhdXRob3JpemF0aW9uIjp7InBlcm1pc3Npb25zIjpbeyJzY29wZXMiOlsib3JkZXI6bGlzdCIsIm9yZGVyOmRldGFpbCJdLCJyc2lkIjoiNzM5Njc3OGUtZTYxYi00ZmU3LWFmOWYtMzY5MTg0OTRkNjc3IiwicnNuYW1lIjoib3JkZXIifV19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiIyZGQxMDEwNS0zYTc3LTQ1MjUtYjJjMS01ZjZjZmIwNzE1NjUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6IkNhayBPYm9iIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiY2Frb2JvYiIsImdpdmVuX25hbWUiOiJDYWsiLCJmYW1pbHlfbmFtZSI6Ik9ib2IiLCJlbWFpbCI6ImJvYmJ5LmJyaWxsaWFuQGdtYWlsLmNvbSJ9.FTN91cHm4JoarE4co6MrMDhdbsdUuELfbcU0rAGYydf-LrImUHsFnbJ6n0efDtar3Wy4VoxAnCFyTl38uhbg5Q7wKpyAs4hZQyyz9owvKKeR7rzGMGB1tAWhB2hObd3qN_YQvmxJqkwZbeanmeqUAmU5qPxAmyk9r2ZUaUou0um0IO5HfLDsPzu0TJlD35CBTO699lH8ggC7GVcutEBVfHnOJMuTmFM5-2ZlXpU_Q5CGs7MFzRVNKpJiCxSJO4vDjh3l5xUoafU4j1FehU0lxXNPg1Oif0IOZZRR-eHZ-oaDnMJ_8aWFMSf6nMX5QrUihl0dHr3cMNQhFVOe8qd1hw
```
Payload: 
```json
{
     "exp": 1640839414,
     "iat": 1640839114,
     "jti": "c1c27012-325f-4f5a-83ed-5290e5211d0b",
     "iss": "http://localhost:9999/auth/realms/dashboard",
     "aud": "dashboard-api",
     "sub": "1dfdb9c2-12e6-4e5f-adf8-69a4ce1f24db",
     "typ": "Bearer",
     "azp": "dashboard-api",
     "session_state": "2dd10105-3a77-4525-b2c1-5f6cfb071565",
     "acr": "1",
     "realm_access": {
       "roles": [
         "order_viewer"
       ]
     },
     "authorization": {
       "permissions": [
         {
           "scopes": [
             "order:list",
             "order:detail"
           ],
           "rsid": "7396778e-e61b-4fe7-af9f-36918494d677",
           "rsname": "order"
         }
       ]
     },
     "scope": "profile email",
     "sid": "2dd10105-3a77-4525-b2c1-5f6cfb071565",
     "email_verified": true,
     "name": "Cak Obob",
     "preferred_username": "cakobob",
     "given_name": "Cak",
     "family_name": "Obob",
     "email": "bobby.brillian@gmail.com"
   }
```

Please read for key `authorization` in `permission scope`, this package is intended to validate that section.

# Example: Setup in Route
* Import package in route setting file.
```go
import keycloakmiddleware "github.com/cakobob/go-keycloak-middleware"
```
* Setup in route
for example we have scope `order:view`, `order:list`, and `order:update`.
```go
app := gin.Default()
api := app.Group("/api")
{
    api.GET("/order", keycloakmiddleware.ValidateJWT([]string{"order:view", "order:list"}), orderListHandler)
    api.PUT("/order/:id", keycloakmiddleware.ValidateJWT([]string{"order:update"}), orderUpdateHandler)
}
```

# License
[MIT License](LICENSE).
