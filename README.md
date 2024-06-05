# GraphQL server test with Golang

Simple server which mirrors the Spotify's REST API into a GraphQL with Golang libraries.

To setup the credentials you need to input the data into config/directory 

```yaml
spotify:
  auth:
    client_id: "client_id"
    client_secret: "client_secret"
    token_endpoint: "token_endpoint"
    auth_endpoint: "auth_endpoint"
    redirect_url: "redirect_url"
    scopes:
      - "scope_1"
      - "scope_2"
  base: "base"
port: 8080
```

You can get the auth details from https://developer.spotify.com/

Scopes used are:
- "playlist-read-private"
- "playlist-read-collaborative"
- "playlist-modify-private"
- "playlist-modify-public"

Depending on what functionalities & entities are used.
