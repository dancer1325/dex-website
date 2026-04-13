# requirements
```
docker run -p 5556:5556 \
  -v $(pwd)/configuration.yaml:/etc/dex/config.yaml \
  dexidp/dex:latest-alpine \
  dex serve /etc/dex/config.yaml
```

# GitHub OAuth2 flow
## if a client redeems a refresh token -- through -- dex -> dex will re-query GitHub /
TODO:
### dex stores a readonly GitHub access token | its backing datastore
TODO:
## if you reject dex's access | GitHub -> revoke ALL dex clients / authenticated them -- through -- GitHub
TODO:

# how to configure?
* | this path
  ```
  docker run -p 5556:5556 \
    -v $(pwd)/configuration.yaml:/etc/dex/config.yaml \
    dexidp/dex:latest-alpine \
    dex serve /etc/dex/config.yaml
  ```
* | Dex's repo
  * `./bin/example-app --issuer http://127.0.0.1:5556/dex --client-id <CLIENT_ID> --client-secret <CLIENT_SECRET>`
    * `--client-id <DEX_CONNECTORS[*].id>`
    * `--client-secret <DEX_CONNECTORS[*].secret>`
* | browser,
  * http://127.0.0.1:5555
    * Authorization code flow > Redirect Github OAuth
      * Problems:
        * Problem1: "request URL http://127.0.0.1:5556/dex/callback?code=a6dbd4bc27a946eebba4&state=il3edfnb3smvxqvncwwds44gc 500 Internal Server Error"
          * Solution: remove | Dex's server configuration `connectors[*].config.orgs`
      * 's return
        * JWT

# Github Enterprise
* TODO:
