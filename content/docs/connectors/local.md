---
linkTitle: "BuiltIn (local)"
title: "Authentication Through the builtin connector"
description: ""
date: 2024-01-05
draft: false
toc: true
weight: 2110
---

* goal
  * Dex's built-in local connector

* Dex's built-in local connector
  * == "virtual" identity provider | Dex's ecosystem
  * securely store login credentials | specified [storage](../configuration/storage)
  * allows
    * simplifying authentication workflows

## how to configure?

* | Dex's configuration file
  * `enablePasswordDB: true`

### how to create users?

* ways
  * [statically | configuration file](#statically--configuration-file)
  * [dynamically -- through the -- gRPC API](#dynamically)

#### statically | configuration file

* | Dex's configuration file
  * `staticPasswords`

#### dynamically

* gRPC API
  * allows
    * handle user-related operations | system

### how to obtain a token?

* | Dex' configuration file
  * `oauth2.passwordConnector`
