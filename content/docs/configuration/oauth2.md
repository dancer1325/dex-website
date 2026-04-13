---
title: "OAuth2"
description: "OAuth2 flow customization options"
date: 2024-01-05
draft: false
toc: true
weight: 1060
---

* goal
  * Dex's configurable options -- about --
    * flow
      * [authentication flow](#authentication-flow)
      * [user flow](#user-flow)
    * grant

## Flow customization

* | Dex's configuration file
  * `oauth2`
* [configuration](https://github.com/dancer1325/dex/blob/master/cmd/dex/config.go)

### Authentication flow

* -- depend on -- `oauth2.responseTypes`

* _Examples:_ [official openid spec](https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationExamples)

### User flow

* allows you to
  * influence how users login | your application

* | Dex's configuration file
  * `oauth2.skipApprovalScreen`
  * `oauth2.alwaysShowLoginScreen`

## Grants customization

* | Dex's configuration file
  * `oauth2.grantTypes`

* [configuration](https://github.com/dancer1325/dex/blob/master/cmd/dex/config.go)

### Available grant types

* [configuration](https://github.com/dancer1325/dex/blob/master/cmd/dex/config.go)

* ALLOWED flows
  * -- depend on -- `oauth2.responseTypes` & `oauth2.grantTypes`

### Default behavior

* default grant types
  * == enabled, by default
  * are
    - `authorization_code`
    - `refresh_token`
    - `urn:ietf:params:oauth:grant-type:token-exchange`
