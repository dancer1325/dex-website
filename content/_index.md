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

## vs OTHER OIDC providers?
### vs Keycloack
* Dex
  * == 1! static Go binary
  * 's backing database
    * pluggable
    * OPTIONAL

* Keycloack
  * == heavy Java application
  * 's backing database
    * MANDATORY

### vs Ory Hydra
* Dex
  * contains
    * login flow (== login UI) + upstream connectors (LDAP, SAML, GitHub, OIDC)

* Hydra
  * == headless OAuth2/OIDC server
    * user is redirected -- to a -- login app / you write
    * login app accept -- , calling Hydra's admin API, --  consent

### vs Authelia / OAuth2 Proxy
* Protocol provider
  * != gateway

* BOTH
  * | HTTP layer /
    * protect -- , via ForwardAuth headers, -- upstream routes

* Dex
  * == FULL OIDC issuer
    * == any standards-compliant client runs the authorization-code flow against it directly

### vs Zitadel
* Dex
  * == Protocol adapter
    * != data plane
    * -> NO owns user state
      * Reason:🧠
        * takes an upstream identity source (LDAP, GitHub, SAML)
        * re-exposes it -- as -- OIDC🧠

* Zitadel
  * == event-sourced IAM platform
    * owns
      * users
      * organisations
      * projects
      * audit logs | CockroachDB

### vs Authentik

* Federates upstream
  * != replace it

* Authentik
  * == Python/Django application /
    * OWN
      * user store
      * flows engine
      * admin UI

* Dex
  * has NO user store
    * authentication is delegated -- to -- whatever IdP / you ALREADY run
  * translates the response -- into -- standard OIDC claims

### vs Cognito / GCP Identity
* Dex
  * runs | your infrastructure
    * -- as a --
      * Kubernetes deployment OR
      * systemd unit OR
      * container
    * -> user data & audit logs stay | your network

* Managed IAM
  * binds your auth path -- to -- cloud's APIs, billing and outage surface


|                             | Dex                             | Keycloak | Ory Hydra        | Authentik     | Zitadel       | Cognito/GCP   |
|-----------------------------|---------------------------------|----------|------------------|---------------|---------------|---------------|
| Runtime                     | Go binary                       | JVM      | Go binary        | Python/Django | Go            | Managed cloud |
| Own user store              | ❌                               | ✅        | ❌                | ✅             | ✅             | ✅             |
| Login UI                    | ✅ minimal                       | ✅ full   | ❌ (you build it) | ✅ full        | ✅ full        | ✅             |
| Upstream connectors         | ✅ (LDAP, SAML, GitHub, OIDC...) | ✅        | ❌                | ✅             | ✅             | limited       |
| Database required           | ❌ pluggable/optional            | ✅        | ✅                | ✅             | ✅ CockroachDB | N/A           |
| OIDC issuer                 | ✅                               | ✅        | ✅                | ✅             | ✅             | ✅             |
| Self-hosted                 | ✅                               | ✅        | ✅                | ✅             | ✅             | ❌             |
| Data stays \|  your network | ✅                               | ✅        | ✅                | ✅             | ✅             | ❌             |
| Complexity                  | low                             | high     | medium           | medium        | high          | low (managed) |
