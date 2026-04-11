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


# Github Enterprise
* TODO:
