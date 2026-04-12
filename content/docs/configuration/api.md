---
title: "The Dex API"
linkTitle: "gRPC API"
description: "Configure Dex dynamically with the gRPC API"
date: 2020-09-30
draft: false
toc: true
weight: 1060
---

* goal
  * how to interact with the Dex's provided API

* Dex's provided [gRPC](http://www.grpc.io/) service
  * allows
    * modifying programmatically the dex's state
  * uses
    * manage applications-- through -- exposing hooks
      * see [API's methods](https://github.com/dexidp/dex/blob/master/api/v2/api.proto)
  * ❌NOT uses❌
    * by MOST installations

## how to configure?

* | Dex's config file
  * `grpc`
    * by default, off

## Clients

* [Dex's gRPC API schema](https://github.com/dexidp/dex/blob/master/api/v2/api.proto)

### Go

* if you want to use -> `go get github.com/dexidp/dex/api/v2`

* _Example:_ [here](https://github.com/dexidp/dex/tree/master/examples/grpc-client/README.md)

### OTHER languages

* steps
  * install [`protoc`](https://github.com/google/protobuf/releases)
  * download the "api.proto" file

## Authentication & access control

* Dex API
  * ⚠️ONLY provide⚠️
    * TLS client auth

* projects / wish to add access controls | EXISTING API
  * should
    * build apps / perform those checks

* _Example:_ if you want to provide a "Change password" screen ->
  * client app
    * could authenticate -- , via Dex's OpenID Connect flow, --  an end user
    * update -- , by calling Dex's API, -- that user's password

## dexctl?

* == Dex CL tool
  * ❌NOT ship -- with -- Dex❌
    * Reason:🧠
      * hard to version
      * easy to design poorly
      * interface / DIFFICULT to change -- due to -- compatibility🧠

## Why NOT REST NOR gRPC Gateway?

* | Dex v1 -> v2,
  * ⚠️REST was migrated -- to -- gRPC⚠️
    * pros (vs [Google APIs](https://github.com/google/apis-client-generator) OR [Open API/Swagger](https://openapis.org/) OR [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway))
      * documentation is generated AUTOMATICALLY -- from -- ".proto"
      * generate AUTOMATICALLY -- , via `protoc`, -- clients / DIFFERENT languages
      * clearly express REST semantics
    * cons
      * MAINLY short term cons

* Dex's configuration file's `staticClients`
  * uses
    * edit clients MANUALLY  | testing
  * ALTERNATIVE to
    * configure the API
