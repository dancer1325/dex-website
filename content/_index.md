---
title: Dex
---
## Why Dex?

### 🔗 Integrate Everything

* steps
  * connect any service -- , through OIDC, to -- Dex
    * 1! time
  * add ANY identity provider (_Examples:_ LDAP, SAML, OIDC)
    * 👀-- through -- connectors👀
    * ⚠️WITHOUT touching your application code⚠️
    * == defer authentication -- to -- OTHER identity providers

### ☸️ Kubernetes Native

* == built -- for -- cloud-native environments
  * -> MINIMAL configuration

* uses
  * SSO | your
    * Kubernetes dashboard
    * internal tools

### ✓ Production Ready

* used by
  * organizations worldwide

* characteristics
  * lightweight
  * standards-based
  * flexible
  * battle-tested
    * Reason:🧠built-in mock provider -- for -- testing | development🧠

### ⚡ Wide Provider Support

* ⚠️requirements⚠️
  * your applications ONLY implement OIDC

* [supported Identity Providers](docs/connectors)

| Provider       | Docs                                       |
|----------------|--------------------------------------------|
| GitHub         | [here](docs/connectors/github.md)          |
| Google         | [here](docs/connectors/google.md)          |
| Microsoft      | [here](docs/connectors/microsoft.md)       |
| LDAP           | [here](docs/connectors/ldap.md)            |
| SAML           | [here](docs/connectors/saml.md)            |
| OpenID Connect | [here](docs/connectors/oidc.md)            |
| GitLab         | [here](docs/connectors/gitlab.md)          |
| LinkedIn       | [here](docs/connectors/linkedin.md)        |
| Atlassian      | [here](docs/connectors/atlassian-crowd.md) |
| Gitea          | [here](docs/connectors/gitea.md)           |
| OAuth 2.0      | [here](docs/connectors/oauth.md)           |
| AuthProxy      | [here](docs/connectors/authproxy.md)       |

## how to use it?

* -- as -- dependency
* [here](docs/guides/using-dex.md)
