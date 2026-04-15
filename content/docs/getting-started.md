---
title: "Getting Started"
description: "First touch with Dex"
date: 2020-09-30
draft: false
toc: true
weight: 1010
type: "docs"
---

## ways to install + run
### -- from -- container image

* == 👀MAIN way👀
* published |
  - [ghcr.io/dexidp/dex](https://github.com/dexidp/dex/pkgs/container/dex)
  - [docker.io/dexidp/dex](https://hub.docker.com/r/dexidp/dex/tags)
* EXISTING variants
  * `alpine`
  * `distroless`
* `dex serve <PATH_TO_CONFIGURATION>`
  * way to run

### Helm chart
* [here](https://github.com/dancer1325/dex-helm)

### -- from -- source code

* requirements
  * install Go v1.19+

* steps
  * | [source repo's host path](https://github.com/dancer1325/dex),
    * `make build`
      * generate "./bin"
    * `./bin/dex serve examples/config-dev.yaml`
  * | browser,
    * http://127.0.0.1:5556/dex
      * check it's running

## Configuration

* 💡-- via -- a config file💡
  * 1! way to configure Dex

* _Example:_ [here](https://github.com/dexidp/dex/blob/master/examples/config-dev.yaml)

### Templated configuration

* | [install -- via -- container images](#---from----container-image),
  * any environment field can be templated | configuration file
    * Reason:🧠default entrypoint
      * pre-process configuration files (".tpl", ".tmpl", ".yaml") -- , via [gomplate](https://github.com/hairyhenderson/gomplate), passed as -- arguments
      * _Example:_ if you use Docker -> `ENTRYPOINT`🧠

## how to run Dex?
* `./bin/dex serve <PATH_TO_CONFIGURATION_FILE>`

## how to run a client?

* ⚠️requirements⚠️
  * [run Dex](#ways-to-install--run)

* Dex's
  * behaviour == MOST OTHER OAuth2 providers' behaviour
    * == | client app | login users,
      * they are redirected -- to -- Dex
  * 👀built-in example client app👀
    * | source repo, built -- via -- `make examples` command
    * uses
      * testing
      * demos
    * 's configuration's OAuth2 credentials
      * == [examples/config-dev.yaml](https://github.com/dexidp/dex/blob/master/examples/config-dev.yaml)'s credentials
    * if you want to run -> `./bin/example-app`
      * query Dex's [discovery endpoint](https://openid.net/specs/openid-connect-discovery-1_0-17.html#ProviderMetadata)
        * == "*/.well-known/openid-configuration"
        * Reason:🧠get the OAuth2 endpoints🧠

* if you want to login | Dex, -- through the -- example app
  * steps
    1. http://localhost:5555/
    2. Click "Authorization Code Flow"
       * Reason:be redirected -- to -- Dex
    3. Choose an option to authenticate:
       - "Login with Example"
         - == use mocked user data
       - "Login with Email"
         - == fill the form -- with -- static user credentials
           - `admin@example.com`
           - `password`
    4. Approve the example app's request
    5. | Dex,
       * see the resulting token
