---
title: "Integration kubelogin and Active Directory"
linkTitle: "Integration kubelogin and Active Directory"
description: ""
date: 2020-09-30
draft: false
toc: true
weight: 2140
---

* goal
  * how dex can work with kubelogin & Active Directory

## Overview

* kubelogin
  * == helper tool / enable integrate kubernetes -- & -- oidc integration
  * uses
    * makes easier to login Open ID Provider

## Requirements

1. Active Directory OR LDAP / has Active Directory compatible schema
   * _Example:_ samba ad
   * enable TLS

2. install [kubelogin](https://github.com/int128/kubelogin/releases)

## Getting started

### Generate certificate & private key

* create OpenSSL conf

  ```bash, title=req.conf
  [req]
  req_extensions = v3_req
  distinguished_name = req_distinguished_name

  [req_distinguished_name]

  [ v3_req ]
  basicConstraints = CA:FALSE
  keyUsage = nonRepudiation, digitalSignature, keyEncipherment
  subjectAltName = @alt_names

  [alt_names]
  # choose your favorite hostname
  DNS.1 = dex.example.com
  ```

* Generate certificate + private key

  ```bash
  $ openssl req -new -x509 -sha256 -days 3650 -newkey rsa:4096 -extensions v3_req -out openid-ca.pem -keyout openid-key.pem -config req.cnf -subj "/CN=kube-ca" -nodes
  $ ls openid*
  openid-ca.pem openid-key.pem
  ```

### Modify dex config

* | Dex config,
  * add

    ```yaml
    connectors:
    - type: ldap
      name: OpenLDAP
      id: ldap
      config:
        host: ldap.example.com:636

        # No TLS for this setup.
        insecureNoSSL: false
        insecureSkipVerify: true

        # This would normally be a read-only user.
        bindDN: cn=Administrator,cn=users,dc=example,dc=com
        bindPW: admin0!
    ```

### Run dex

```bash
$ bin/dex serve examples/config-ad-kubelogin.yaml
```

### Configure kubernetes with oidc

TODO:
Copy `openid-ca.pem` to `/etc/ssl/certs/openid-ca.pem` on master node.

Use the following flags to point your API server(s) at dex
`dex.example.com` should be replaced by whatever DNS name or IP address dex is running under.

```bash
--oidc-issuer-url=https://dex.example.com:32000/dex
--oidc-client-id=kubernetes
--oidc-ca-file=/etc/ssl/certs/openid-ca.pem
--oidc-username-claim=email
--oidc-groups-claim=groups
```

Then restart API server(s).


See https://kubernetes.io/docs/reference/access-authn-authz/authentication/ for more detail.

### Set up kubeconfig

Add a new user to the kubeconfig for dex authentication:

```bash
$ kubectl config set-credentials oidc \
    --exec-api-version=client.authentication.k8s.io/v1beta1 \
    --exec-command=kubectl \
    --exec-arg=oidc-login \
    --exec-arg=get-token \
    --exec-arg=--oidc-issuer-url=https://dex.example.com:32000/dex \
    --exec-arg=--oidc-client-id=kubernetes \
    --exec-arg=--oidc-client-secret=ZXhhbXBsZS1hcHAtc2VjcmV0 \
    --exec-arg=--oidc-extra-scope=profile \
    --exec-arg=--oidc-extra-scope=email \
    --exec-arg=--oidc-extra-scope=groups \
    --exec-arg=--certificate-authority-data=$(base64 -w 0 openid-ca.pem)
```

Please confirm `--oidc-issuer-url`, `--oidc-client-id`, `--oidc-client-secret` and `--certificate-authority-data` are same as values in config-ad-kubelogin.yaml.

Run the following command:

```bash
$ kubectl --user=oidc cluster-info
```

It launches the browser and navigates it to http://localhost:8000.
Please log in with your AD account (eg. test@example.com) and password.
After login and grant, you can access the cluster.

You can switch the current context to dex authentication.

```bash
$ kubectl config set-context --current --user=oidc
```
