<pre>
server
    {
        listen 80;
        #listen [::]:80;
        server_name go.kitcloud.cn;
        location / {

                proxy_redirect off;
                proxy_set_header Host "golang.org";
                proxy_pass https://golang.org;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                access_log off;
                error_log off;
        }
    }
</pre>