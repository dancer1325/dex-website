
* public client

  ```shell
  curl -L -X POST 'http://localhost:8080/dex/token' \
  # Authorization: Basic base64(clientId)
  -H 'Authorization: Basic cHVibGljLWNsaWVudDo=' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'grant_type=password' \
  --data-urlencode 'scope=openid profile' \
  --data-urlencode 'username=admin@example.com' \
  --data-urlencode 'password=password'
  ```

* private client

  ```shell
  curl -L -X POST 'http://localhost:8080/dex/token' \
  # Authorization: Basic base64(clientId:clientSecret)
  -H 'Authorization: Basic cHJpdmF0ZS1jbGllbnQ6YXBwLXNlY3JldA=='
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'grant_type=password' \
  --data-urlencode 'scope=openid' \
  --data-urlencode 'username=admin@example.com' \
  --data-urlencode 'password=password'
  ```
