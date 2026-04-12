---
title: "Storage"
description: "Configuration options for persistent data storage"
date: 2020-09-30
draft: false
toc: true
weight: 1050
---

* goal
  * Dex' supported storage configurations


* Dex persist state
  * ⚠️requirements⚠️
    * transport security
    * database ACLs
  * Reason:🧠
    * Dex perform MULTIPLE tasks
      * _Examples:_ track refresh tokens, preventing replays, and rotating keys
    * OTHERWISE, they can affect applications / rely on dex🧠
  * _Example:_ sensitive data (signing keys, bcrypt'd passwords)

## Etcd

* == persist state | [etcd v3](https://github.com/coreos/etcd)

* [configuration](https://github.com/dancer1325/dex/blob/master/storage/etcd/config.go)

## Kubernetes CRDs

* requirements
  * Kubernetes v1.7+
  * Dex
    * can access -- to the -- NON-namespaced CRD type
      * == set `ClusterRole` & `ClusterRoleBinding`
      * recommendation: Dex run | Kubernetes cluster
      * way / Dex determines the namespace | it's running
        * parse AUTOMATICALLY the service account token / mounted | its pod

* replacement of
  * Kubernetes Third Party Resource (TPR) extension

* CRDs
  * enable
    * Dex can run | EXISTING Kubernetes cluster WITHOUT an external database

* `dex.coreos.com.AuthCode`
  * use cases
    * debugging
  * uses
    * store state
  * ❌NOT uses❌
    * by end users

### Kubernetes third party resources (TPRs)

* deprecated
  * Dex [v2.17.0-](https://github.com/dexidp/dex/tree/v2.17.0)

TODO:
If you are currently running dex using TPRs, you will need to [migrate to CRDs](https://github.com/dexidp/dex/blob/v2.17.0/Documentation/storage.md#migrating-from-tprs-to-crds)
before you upgrade to a post v2.17 dex
*  The script mentioned in the instructions can be [found here](https://github.com/dexidp/dex/blob/v2.17.0/scripts/dump-tprs)

### how to configure?

* | Dex configuration file
  * `storage.type: kubernetes`

## SQL

* supported flavors of SQL
  * [SQLite3](#sqlite3)
  * [Postgres](#postgres)
  * [MySQL](#mysql)

* Migrations
  * == database schema changes
  * | FIRST connection -- to the -- SQL server, are performed AUTOMATICALLY
    * == ❌NOT support rolling back❌
    * requirements
      * Dex need requires privileges -- to -- add & alter the database's tables

* ⚠️requirements⚠️
  * | PREVIOUS Dex versions,
    * symmetric keys
      * Reason: 🧠BEFORE sending values to the database -> encrypt them🧠
      * ❌NOT ported -- to -- Dex v2❌
    * TODO: If it is added later there may not be a migration path for current v2 users.

### SQLite3

* use cases
  * stand up dex QUICKLY
* ❌NOT use cases❌
  * real workloads

* [configuration](https://github.com/dancer1325/dex/blob/master/storage/sql/sqlite.go)

* if the `:memory:` value is provided | SQLite3 -> dex AUTOMATICALLY disable support -- for -- CONCURRENT database queries
  * Reason:🧠SQLite3 prevent race conditions -- via -- file locks🧠

### Postgres

* recommendations
  * add a DEDICATE a database -- to -- Dex
    * Reason:🧠
      1. Dex needs privileged access -- to -- its database
         * Reason: 🧠it performs migrations🧠
      2. | share with OTHER applications, Dex's database table names are NOT configurable🧠

* [configuration](https://github.com/dancer1325/dex/blob/master/storage/ent/postgres.go)

### MySQL

* requirements
  * MySQL v5.7+

* recommendations
  * add a DEDICATE a database -- to -- Dex
    * Reason:🧠
      1. Dex needs privileged access -- to -- its database
        * Reason: 🧠it performs migrations🧠
      2. | share with OTHER applications, Dex's database table names are NOT configurable🧠

* [configuration](https://github.com/dancer1325/dex/blob/master/storage/ent/mysql.go)

## how to add a NEW storage options?

TODO:
Each storage implementation bears a large ongoing maintenance cost and needs to be updated every time a feature requires storing a new type
* Bugs often require in depth knowledge of the backing software, and much of this work will be done by developers who are not the original author
* Changes to dex which add new storage implementations require a strong use case to be considered for inclusion.

### New storage option references

Those who still want to construct a proposal for a new storage should review the following packages:

* `github.com/dexidp/dex/storage`: Interface definitions which the storage must implement
  * __NOTE:__ This package is not stable.
* `github.com/dexidp/dex/storage/conformance`: Conformance tests which storage implementations must pass.

### New storage option requirements

Any proposal to add a new implementation must address the following:

* Integration testing setups (Travis and developer workstations).
* Transactional requirements: atomic deletes, updates, etc.
* Is there an established and reasonable Go client?

[issues-transaction-tests]: https://github.com/dexidp/dex/issues/600
[k8s-api]: https://github.com/kubernetes/kubernetes/blob/master/docs/devel/api-conventions.md#concurrency-control-and-consistency
[psql-conn-options]:
[mysql-conn-options]: https://github.com/go-sql-driver/mysql#tls
