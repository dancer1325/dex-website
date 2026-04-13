---
title: "Tokens"
linkTitle: "Tokens"
description: "Types of tokens and expiration settings"
date: 2020-10-21
draft: false
toc: true
weight: 1013
---

* ID Tokens
  * == 💡OAuth2 extension💡 /
    * introduced -- by -- OpenID Connect
    * 👀Dex's primary feature👀
  * == 👀[JSON Web Tokens](https://jwt.io/)👀 /
    * signed -- by -- Dex
    * ⚠️contained | OAuth2 response⚠️
    * contain [standard claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)
    * attest the end user's identity
  * == `base64First.base64Second.base64Third`
    * `base64First`
      * == header
    * `base64Second`
      * == payload
    * `base64Third`
      * == signature of (`base64First`, `base64Second`)
  * _Example:_

    ```bash
    eyJhbGciOiJSUzI1NiIsImtpZCI6IjlkNDQ3NDFmNzczYjkzOGNmNjVkZDMyNjY4NWI4NjE4MGMzMjRkOTkifQ.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjU1NTYvZGV4Iiwic3ViIjoiQ2djeU16UXlOelE1RWdabmFYUm9kV0kiLCJhdWQiOiJleGFtcGxlLWFwcCIsImV4cCI6MTQ5Mjg4MjA0MiwiaWF0IjoxNDkyNzk1NjQyLCJhdF9oYXNoIjoiYmk5NmdPWFpTaHZsV1l0YWw5RXFpdyIsImVtYWlsIjoiZXJpYy5jaGlhbmdAY29yZW9zLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJncm91cHMiOlsiYWRtaW5zIiwiZGV2ZWxvcGVycyJdLCJuYW1lIjoiRXJpYyBDaGlhbmcifQ.OhROPq_0eP-zsQRjg87KZ4wGkjiQGnTi5QuG877AdJDb3R2ZCOk2Vkf5SdP8cPyb3VMqL32G4hLDayniiv8f1_ZXAde0sKrayfQ10XAXFgZl_P1yilkLdknxn6nbhDRVllpWcB12ki9vmAxklAr0B1C4kr5nI3-BZLrFcUR5sQbxwJj4oW1OuG6jJCNGHXGNTBTNEaM28eD-9nhfBeuBTzzO7BKwPsojjj4C9ogU4JQhGvm_l4yfVi0boSx8c0FX3JsiB0yLa1ZdJVWVl9m90XmbWRSD85pNDQHcWZP9hR6CMgbvGkZsgjG32qeRwUL_eNkNowSBNWLrGNPoON1gMg
    ```
    * unconverting it

      ```json
      {
      "iss": "http://127.0.0.1:5556/dex",
      "sub": "CgcyMzQyNzQ5EgZnaXRodWI",
      "aud": "example-app",
      "exp": 1492882042,
      "iat": 1492795642,
      "at_hash": "bi96gOXZShvlWYtal9Eqiw",
      "email": "jane.doe@coreos.com",
      "email_verified": true,
      "groups": [
      "admins",
      "developers"
      ],
      "name": "Jane Doe"
      }
      ```
      * `iss`
        * == server / issued this token
      * `sub`
        * == token's subject
        * == end user's UNIQUE ID
      * `aud`
        * == token's audience
          * == OAuth2 client ID / this was issued for

  * uses
    * by other services -- as -- service-to-service credentials
      * supported services
        * [Kubernetes](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens)
        * [AWS STS](https://docs.aws.amazon.com/STS/latest/APIReference/Welcome.html)
    * [MORE](../guides/using-dex.md)

## Refresh tokens

* Refresh tokens
  * == credentials /
    * uses
      * obtain NEW id tokens | CURRENT id token becomes invalid OR expires
    * are rotated
      * == [refresh_token rotation](https://tools.ietf.org/html/rfc6819#section-5.2.2.3)
      * == / EACH id token refresh, issue NEW refresh token
        * -- to the -- client
        * -- by the -- authorization server
      * allows
        * prevents someone stealing it
  * ⚠️OPTIONAL⚠️
    * if you want it -> passes `offline_access` scope -- to -- Dex server
  * _Example of server response / contains refresh token:_

    ```json
    {
     "access_token": "eyJhbGciOiJSUzI1N...",
     "token_type": "Bearer",
     "refresh_token": "lxzzsvasxho5exvwkfa5zhefl",
     "expires_in": 3600,
     "id_token": "eyJhbGciO..."
    }
    ```

### how to configure?

* | Dex's configuration file
  * `.expiry`

## Token signing configuration

* | Dex's configuration file
  * `.signer`
    * [interface](https://github.com/dancer1325/dex/blob/master/server/signer/signer.go)

* ALLOWED types
  * [local signer](#local-signer)
  * [Vault-compatible signer](#vault-compatible-signer)

### Local signer

* [configuration](https://github.com/dancer1325/dex/blob/master/server/signer/local.go)
* use cases
  * simple deployments

### Vault-compatible signer

* supported -- through -- [OpenBao API v2 integration package](https://pkg.go.dev/github.com/openbao/openbao/api/v2)
* [configuration](https://github.com/dancer1325/dex/blob/master/server/signer/vault.go)
* allows
  * signing operations -- through -- HashiCorp Vault OR OpenBao
    * WITHOUT storing keys LOCALLY
