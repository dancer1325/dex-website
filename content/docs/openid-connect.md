---
title: "An Overview of OpenID Connect"
linkTitle: "Intro to OpenID Connect"
description: "Intro to OpenID Connect (basics)"
date: 2020-09-30
draft: false
toc: true
type: "docs"
weight: 1020
---

* goal
  * [OpenID Connect protocol](https://openid.net/connect/)

* OpenID Connect protocol
  * == flavor of OAuth2 /
    * Dex implements

## OAuth2

* uses
  * "Login with Google" button

* | server side apps,
  * general flow
    1. NEW user visits an application
    2. application redirects the user -- to -- Dex
    3. user logs (-- via -- connectors) | Dex
       * is asked if it's okay to let the application
         * view the user's profile
         * post on their behalf
         * etc.
    4. if the user clicks okay -> Dex redirects the user -- , with a code, back to the -- application
    5. the application redeems that code with provider -- for a -- token

## ID Tokens

* access token applications / get from OAuth2 providers
  * completely opaque -- to the -- client
  * UNIQUE -- to the -- provider
    * == token / received from Google != token / received from Twitter OR GitHub OR ...

* _Example of token response -- from an -- OpenID Connect:_

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json
  Cache-Control: no-store
  Pragma: no-cache

  {
   "access_token": "SlAV32hkKG",
   "token_type": "Bearer",
   "refresh_token": "8xLOxBtZp8",
   "expires_in": 3600,
   "id_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFlOWdkazcifQ.ewogImlzcyI6ICJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwKICJzdWIiOiAiMjQ4Mjg5NzYxMDAxIiwKICJhdWQiOiAiczZCaGRSa3F0MyIsCiAibm9uY2UiOiAibi0wUzZfV3pBMk1qIiwKICJleHAiOiAxMzExMjgxOTcwLAogImlhdCI6IDEzMTEyODA5NzAKfQ.ggW8hZ1EuVLuxNuuIJKX_V8a_OMXzR0EHR9R6jgdqrOOF4daGU96Sr_P6qJp6IcmD3HP99Obi1PRs-cwh3LO-p146waJ8IhehcwL7F09JdijmBqkvPeB2T9CJNqeGpe-gccMg4vfKjkM8FcGvnzZUN4_KSP0aAp1tOJ1zZwgjxqGByKHiOtX7TpdQyHE5lcMiKPXfEIQILVq0pc_E2DzL7emopWoaoZTF_m0_N0YzFC6g6EJbOEoRoSK5hoDalrcvRYLSrQAZZKflyuVCyixEoV9GfNQC3_osjzw2PAithfubEEBLuVVk4XUVrWOLrLl0nx7RkKU8NXNHq-rvKMzqg"
  }
  ```

* [ID tokens](configuration/tokens.md)

## Discovery

* OpenID Connect servers' discovery mechanism -- "/.well-known/openid-configuration" --
  * 's response == OpenID Connect features
    * OAuth2 endpoints
      * uses
        * OAuth2 clients
    * token endpoint
      * uses
        * OAuth2 clients
    * scopes supported
    * ...

* "/keys"
  * 's return
    * [JSON Web Key](https://tools.ietf.org/html/rfc7517)
