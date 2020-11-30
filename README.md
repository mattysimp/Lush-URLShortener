# Lush-URLShortener

The URLShortener will take a provided url in a json format in a POST request and create a short url which will be created as "base-url/GUID" and returned to the requester. It will also accept a generated shorturl and redirect the user to the long url.

Short URLs will be an 32 bit sequence of case-sensitive Alphanumeric characters generated through the CRC32 hashing algorithm with the addition of amendment characters which will increment up in the case of hash collisions. The amendment character starts at 0.

### Storage
All URLs will be held in a dockerised MongoDB. 

### Examples
```curl -X POST -d '{"url":"https://uk.lush.com/christmas/bath-bombs/i-want-a-hippopotamus-for-christmas"}' localhost:8080 -D -
HTTP/1.1 200 OK
Content-Type: application/json

{"short_url":"localhost:8080/aB8uISc90"}```


```curl -X GET -d '' localhost:8080/aB8uISc90 -D -
HTTP/1.1 200 OK
Content-Type: application/json

{"url":"https://uk.lush.com/christmas/bath-bombs/i-want-a-hippopotamus-for-christmas", "short_url":"localhost:8080/aB8uISc90", "url_code":"aB8uISc90"}```