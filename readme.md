# Test available urls (as TAU)

This is a simple utility that checks the responses of a list of URLs.<br/>
For simplicity, the URLs are taken from the settings.json configuration file.

```
{
    "urls": [
       {"address": "127.0.0.1"},
       {"address": "localhost"},
       {"address": "127.0.0.2"}
    ],
    "options": {
        "ssl_mode": "http",
        "http_timeout": 1
    }
}
```

If the URL did not respond, the report will have a code of 0. <br/>
The result of the utility in a unique generated file ``report_<request_datetime>.json``