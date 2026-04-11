---
title: "Connectors"
description: "Documentation about configuration of Dex connectors"
date: 2020-01-07T14:59:38+01:00
draft: false
toc: true
weight: 2000
---

* [architecture](../_index.md)

![](../../../static/img/dex-flow.png)

* clients
  * SOME
    * need refresh token functionality
      * POSSIBLE Reason: 🧠require OFFLINE access🧠
      * _Example:_ `kubectl`

* connector
  * == Dex's strategy /
    * enable: user can authenticate -- against --
      * ANOTHER protocols
        * _Example: LDAP, SAML
      * specific platforms
        * _Examples:_ GitHub, LinkedIn, and Microsoft
  * ⚠️limitations⚠️
    * -- depending on -- EACH protocol
      * issue [refresh tokens](../configuration/custom-scopes-claims-clients.md#scopes)
        * _Example:_ SAML connector
          * Reason:🧠SAML ONLY refresh tokens -- through -- browser redirections + user interaction🧠
      * return [group membership](../configuration/custom-scopes-claims-clients.md#scopes) claims
  * implemented connectors

| Name                                                 | supports refresh tokens | supports groups claim | supports preferred_username claim | status  | notes |
|------------------------------------------------------|-------------------------|-----------------------|-----------------------------------|---------| ----- |
| [LDAP](/docs/connectors/ldap/)                       | yes                     | yes                   | yes                               | stable  | |
| [GitHub](/docs/connectors/github/)                   | yes                     | yes                   | yes                               | stable  | |
| [SAML 2.0](/docs/connectors/saml/)                   | no                      | yes                   | no                                | stable  |
| [GitLab](/docs/connectors/gitlab/)                   | yes                     | yes                   | yes                               | beta    | |
| [OpenID Connect](/docs/connectors/oidc/)             | yes                     | yes                   | yes                               | beta    | Includes Salesforce, Azure, etc. |
| [OAuth 2.0](/docs/connectors/oauth/)                 | no                      | yes                   | yes                               | alpha   |
| [Google](/docs/connectors/google/)                   | yes                     | yes                   | yes                               | alpha   | |
| [LinkedIn](/docs/connectors/linkedin/)               | yes                     | no                    | no                                | beta    | |
| [Microsoft](/docs/connectors/microsoft/)             | yes                     | yes                   | no                                | beta    | |
| [AuthProxy](/docs/connectors/authproxy/)             | no                      | no                    | no                                | alpha   | Authentication proxies such as Apache2 mod_auth, etc. |
| [Bitbucket Cloud](/docs/connectors/bitbucketcloud/)  | yes                     | yes                   | no                                | alpha   | |
| [OpenShift](/docs/connectors/openshift/)             | no                      | yes                   | no                                | stable  | |
| [Atlassian Crowd](/docs/connectors/atlassian-crowd/) | yes                     | yes                   | yes *                             | beta    | preferred_username claim must be configured through config |
| [Gitea](/docs/connectors/gitea/)                     | yes                     | no                    | yes                               | alpha   | |
| [OpenStack Keystone](/docs/connectors/keystone/)     | yes                     | yes                   | no                                | alpha   |  |

* status
  * Stable
    * ==
      * well tested
      * in active use
      * will NOT change -- as -- backward incompatible ways
  * Beta
    * ==
      * tested
      * unlikely to change -- as -- backward incompatible ways.
  * Alpha
    * ==
      * may be untested -- by -- core maintainers
      * subject to change -- as -- backward incompatible ways
