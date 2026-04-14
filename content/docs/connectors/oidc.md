---
title: "Authentication Through an OpenID Connect Provider"
linkTitle: "OpenID Connect"
description: ""
date: 2020-09-30
draft: false
toc: true
weight: 2050
---

* goal
  * ANOTHER OpenID Connect provider -- as an -- authentication source

* [OpenID Connect protocol](../openid-connect.md)

* _Examples of OpenID Connect providers:_
  * Google Accounts
  * Salesforce
  * Azure AD v2 ([❌NOT v1❌](https://github.com/coreos/go-oidc/issues/133))

## how to configure?

```yaml
connectors:
- type: oidc
  id: google
  name: Google
  config:
    issuer: https://accounts.google.com

    # issuerAlias: https://accounts.google.com

    # $<ENVIRONMENT_VARIABLE>
    clientID: $GOOGLE_CLIENT_ID
    clientSecret: $GOOGLE_CLIENT_SECRET

    # Dex's issuer URL + "/callback"
    redirectURI: http://127.0.0.1:5556/dex/callback

    # basicAuthUnsupported: true

    # scopes:
    #  - profile
    #  - email
    #  - groups

    # insecureSkipEmailVerified: true

    # insecureEnableGroups: true

    # allowedGroups:
    #  - <value>

    # getUserInfo: true

    # userIDKey: nickname

    # userNameKey: nickname

    # acrValues:
    #  - <value>
    #  - <value>

    # promptType: consent

    claimMapping:
      # preferred_username: other_user_name

      # email: mail

      # groups: "cognito:groups"

    claimModifications:
      # _Example of this:_ resulting group == `example::organization::email`
      # newGroupFromClaims:
      #   - prefix: example
      #     delimiter: "::"
      #     clearDelimiter: false
      #     claims:
      #       - organization
      #       - email

      # filterGroupClaims:
      #   groupsFilter: "<REGEX>"

      # _Example:_ if the connector provides a group "regular-users" -> modification converts it -- to -- "example-prefix-regular-usersexample-suffix"
      # modifyGroupNames:
      #   prefix: example-prefix- # note the delimiter at the end
      #   suffix: example-suffix


    overrideClaimMapping: false

    providerDiscoveryOverrides:
      # tokenURL: ""
      # authURL: ""
```
