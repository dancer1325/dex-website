# ways to install + run
## -- from -- container image
### published |
#### https://github.com/dexidp/dex/pkgs/container/dex
TODO:
#### https://hub.docker.com/r/dexidp/dex/tags
* `docker pull dexidp/dex:latest-alpine`
* run it
```
docker run -p 5556:5556 \
  -v $(pwd)/configTwo.yaml:/etc/dex/config.yaml \
  dexidp/dex:latest-alpine \
  dex serve /etc/dex/config.yaml
```

## Helm chart
* [here](https://github.com/dancer1325/dex-helm)

## -- from -- source code
* follow the steps
* check
  * ".bin/" is created
  * `./bin/dex version`
    * check it's installed

# Configuration
## Templated configuration
### | install -- via -- container image,
#### any environment field can be templated | configuration file
* pass the templated inputs
  ```
  docker run -p 5556:5556 \
    -e DEX_ISSUER=http://127.0.0.1:5556/dex \
    -e DEX_CLIENT_SECRET=mysecret \
    -e DEX_REDIRECT_URI=http://127.0.0.1:5555/callback \
    -e DEX_USER_EMAIL=admin@example.com \
    -e DEX_USER_PASSWORD_HASH='$2a$10$2b2cU8CPhOTaGrs1HRQuAueS7JTT5ZHsHSzYiFPm1leZck7Mc8T4W' \
    -v $(pwd)/config.yaml:/etc/dex/config.yaml \
    dexidp/dex:latest-alpine \
    dex serve /etc/dex/config.yaml
  ```

# how to run a client?
## requirements: run Dex
```
docker run -p 5556:5556 \
  -v $(pwd)/configTwo.yaml:/etc/dex/config.yaml \
  dexidp/dex:latest-alpine \
  dex serve /etc/dex/config.yaml
```
## Dex's behaviour == MOST OTHER OAuth2 providers' behaviour
### == | client app | login users, they are redirected -- to -- Dex
* TODO:
## built-in example client app
### | source repo, built -- via -- `make examples` command
* | source repo, `make examples`
  * check the logs
### 's configuration's OAuth2 credentials
TODO:
### if you want to run -> `./bin/example-app`
* | source code repo, `./bin/example-app`


# TODO:
