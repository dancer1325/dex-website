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

### Examples

**Enable only Authorization Code flow:**
```yaml
oauth2:
  grantTypes: [ "authorization_code" ]
```

**Enable client credentials grant for server-to-server authentication:**

Set the required environment variable, client credentials grant is enabled by default:
```bash
export DEX_CLIENT_CREDENTIAL_GRANT_ENABLED_BY_DEFAULT=true
```

**Enable password grant (not recommended):**
```yaml
oauth2:
  passwordConnector: local  # Required for password grant
```

Password grants involve clients directly sending a user's credentials (`username` and `password`) to the authorization server (dex), acquiring access tokens without the need for an intermediate authorization step.

**Enable Implicit Flow:**

Implicit flow is configured via `responseTypes`, not `grantTypes`:
```yaml
oauth2:
  responseTypes: [ "id_token", "token" ]
```

### Configuration options

* `grantTypes` - list of enabled grant types (see [Configurable Grants](#configurable-grants) section above). To enable password grants, ensure `"password"` is included in this list.
* `passwordConnector` - specifies the connector's id that is used for password grants

{{% alert title="Warning" color="warning" %}}
The password grant type is not recommended for use by the [OAuth 2.0 Security Best Current Practice](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-security-topics-13#section-3.4) because of serious security concerns.
Please see [oauth.net](https://oauth.net/2/grant-types/password/) for additional information.
{{% /alert %}}
