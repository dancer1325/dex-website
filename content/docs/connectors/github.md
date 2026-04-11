---
title: "Authentication Through GitHub"
linkTitle: "GitHub"
description: ""
date: 2020-09-30
draft: false
toc: true
weight: 2020
---

* goal
  * identify the end user -- through -- their GitHub account

* GitHub OAuth2 flow
  * if a client redeems a refresh token -- through -- dex -> dex will re-query GitHub /
    * dex stores a readonly GitHub access token | its backing datastore
  * if you reject dex's access | GitHub -> revoke ALL dex clients / authenticated them -- through -- GitHub

```mermaid
sequenceDiagram
    participant Client
    participant Dex
    participant GitHub

    Note over Client,GitHub: Refresh Token Flow
    Client->>Dex: refresh token
    Dex->>GitHub: re-query (stored readonly access token)
    GitHub-->>Dex: updated user info
    Dex-->>Client: new ID Token

    Note over Client,GitHub: Revocation Flow
    GitHub->>Dex: access token invalidated (user revoked)
    Dex->>Client: refresh denied
    Dex->>Client2: refresh denied
    Note over Client,Client2: ALL clients authenticated via GitHub lose access
```

## how to configure?

* | [Github](https://github.com/settings/applications/new)
  * == Settings > Developer Settings > OAuth Apps > New OAuth App
  * register a NEW application /
    * callback URL == `(dex issuer)/callback`
      * _Example:_ if dex is listening | NON-root path https://auth.example.com/dex -> the callback == https://auth.example.com/dex/callback

## GitHub Enterprise

### how to generate TLS assets?

TODO:
Running Dex with HTTPS enabled requires a valid SSL certificate, and the API server needs to trust the certificate of the signing CA using the `--oidc-ca-file` flag.

For our example use case, the TLS assets can be created using the following command:

```bash
$ ./examples/k8s/gencert.sh
```

This will generate several files under the `ssl` directory, the important ones being `cert.pem` ,`key.pem` and `ca.pem`
* The generated SSL certificate is for 'dex.example.com', although you could change this by editing `gencert.sh` if required.

### how to run example client app -- with -- GitHub config?

* steps
  * `./bin/example-app --issuer-root-ca examples/k8s/ssl/ca.pem`
  * | browser,
    * http://127.0.0.1:5555
    * click Login
    * select
      * Log in -- via -- GitHub
      * grant access -- to -- dex to view your profile
