# Install

~~~
go install github.com/pkg6/ssl-certificate/ssl-certificate@latest
~~~

# Config

~~~
{
  "domains": [
    {
      "deploy": "ssh",
      "certificate": {
        "domain": [
          "www.zhiqiang.wang"
        ],
        "provider": {
          "name": "aliyun",
          "config": {
            "accessKeyId": "",
            "accessKeySecret": ""
          }
        },
        "registration": {
          "provider": "letsencrypt"
        }
      }
    }
  ],
  "deploys": {
    "ssh": {
      "host": "127.0.0.1",
      "port": 22,
      "username": "ubuntu",
      "password": "123456",
      "beforeCommand": "",
      "afterCommand": "service nginx restart",
      "certPath": "/etc/nginx/ssl/ssh.cer",
      "keyPath": "/etc/nginx/ssl/ssh.key"
    },
    "local": {
      "beforeCommand": "",
      "afterCommand": "",
      "certPath": "/etc/nginx/ssl/local.cer",
      "keyPath": "/etc/nginx/ssl/local.key"
    }
  }
}
~~~

