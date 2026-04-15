# requirements
* [install Dex](../../../getting-started.md)
* [run Dex](../../../getting-started.md)
```
docker run -p 5556:5556 \
  -v $(pwd)/dexConfiguration.yaml:/etc/dex/config.yaml \
  dexidp/dex:latest-alpine \
  dex serve /etc/dex/config.yaml
```
* [run Dex's built-in client app](../../../getting-started.md#how-to-run-a-client)
  * | Dex's repo,
    * `./bin/example-app`

# how to configure?
## _Example:_ Google Account
* requirements
  * [| your google-cloud, create your Google Client Id & Google CLient Secret](https://developers.google.com/identity/protocols/oauth2?hl=es-419)
* steps
  * | [here](dexConfiguration.yaml),
    * specify YOUR <GOOGLE_CLIENT_ID> & <GOOGLE_CLIENT_SECRET>
* test -- through -- Dex's built-in client app
  * http://127.0.0.1:5555/
    * Authorization Code Flow > Login in with Google > Login with your gmail account
    * return JWT
