# TODO:

# Cross-client trust and authorized party
* TODO: [configuration.yaml](configuration.yaml)
* web app can use the following scope -- to -- request an ID token / is issued -- for the -- CL tool

  ```
  audience:server:client_id:cli-app
  ```

* ID token claims include audience + authorized party

  ```json
  {
      "aud": "cli-app",
      "azp": "web-app",
      "email": "foo@bar.com",
      // other claims...
  }
  ```
