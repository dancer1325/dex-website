---
title: "Scopes and Claims"
description: "Custom Scopes, Claims and Client Features"
date: 2020-09-30
draft: false
toc: true
weight: 1040
---

* goal
  * OAuth2's features & OpenID Connect's features /
    * implemented -- by -- dex

## Scopes

* Dex's supported scopes

| Name                                      | Description                                                                                                                                                                                    |
|-------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `openid`                                  | REQUIRED -- for -- ALL login requests                                                                                                                                                          |
| `email`                                   | == end user's email and if that email was verified by an upstream provider                                                                                                                     |
| `profile`                                 | == end user username                                                                                                                                                                           |
| `groups`                                  | groupS / end user is a member of                                                                                                                                                               |
| `federated:id`                            | ID provider information <br/> == connector ID + user ID / assigned \| provider                                                                                                                 |
| `offline_access`                          | \| token response, should contain a refresh token <br/> \| SOME connectors, <br/> &nbsp;&nbsp; does NOT work <br/> &nbsp;&nbsp; it's ignored (_Example:_ [SAML connector](../connectors/saml)) |
| `audience:server:client_id:( client-id )` | == dynamic scope / ID token should be issued -- on behalf of -- ANOTHER client <br/> [here](#cross-client-trust-and-authorized-party)                                                          |

## Custom claims

* == NON-standard claims / implemented -- by -- Dex

* == ADDITIONAL to
  * [required OpenID Connect REQUIRED claims](https://openid.net/specs/openid-connect-core-1_0.html#IDToken) &
  * [standard claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)

| Name                 | Description                                               |
|----------------------|-----------------------------------------------------------|
| `groups`             | `string[]` <br/> groups / a user is a member of           |
| `federated_claims`   | connector ID + user ID / assigned to the user \| provider |
| `email`              | email of the user                                         |
| `email_verified`     | if the upstream provider has verified the email           |
| `name`               | user's display name                                       |
| `preferred_username` | shorthand name / end-user wishes to be referred to        |

* _Example of `federated_claims` claim:_

  ```json
  "federated_claims": {
    "connector_id": "github",
    "user_id": "110272483197731336751"
  }
  ```

## Cross-client trust and authorized party

* Dex
  * can issue ID tokens -- , on behalf of OTHER clients, to -- clients
    * == (| OpenID Connect terms),
      * ID token's `aud` (audience) claim != client / performed the login

## Public clients

* inspired -- by -- [Google's "Installed Applications"](https://developers.google.com/api-client-library/python/auth/installed-app)
* uses
  * impose restrictions | applications /
    * do NOT intend to keep their client secret private

* how to configure?
  * `staticClients[*].public: true`
  * if `staticClients[*].redirectURIs` are NOT specified -> | public clients, default values
    * redirects / begin with "http://localhost"
    * "urn:ietf:wg:oauth:2.0:oob"
      * == special "out-of-browser" URL
      * triggers dex / display the OAuth2 code | browser
      * client's responsibility
        * create a screen OR prompt -- to -- receive the code
        * perform a code exchange -- for a -- token response

* "out-of-browser" flow
  * recommendations
    * ID Token nonce
