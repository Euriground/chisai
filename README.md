# Chisai
 this is a simple url shortening app developed using golang

## How to Use
- fork or clone the repository
- open the app folder 
- run command below

```
    docker-compose up -d --build
```

 ## Endpoints

-  **POST** \</shorten\> <br/>
        send request with body of this format:<br/>
        ```
            {
             "long_url" : "YOUR_WEBSITE_URL"
            }
        ```
 
- **GET** \</:short_url\><br/>
    just put the short value you get after sending POST request here<br/>
    If you're in the browser, you will be redirected to your desired website
