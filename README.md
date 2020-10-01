# Overview
<p> Technical task for Golang dev position at Ciklum. <br>
    Project consists of <code>main</code> package and <code>modules</code> package. <br>
    API has only one path - root path.</p>
    
## Running
### Without Docker
<pre><code>go build CiklumGolangTask/main
go run CiklumGolangTask/main
</code></pre>
### With Docker
<pre><code>docker build -t golang/api_task .
docker run -p 8888:8888 --rm golang/api_task
</code></pre>

## Calling API example
### GET request
<pre><code>GET / HTTP/1.1 
Host: 127.0.0.1:8888
</code></pre>
### URLs table
<table>
    <thead>
        <tr>
            <th>Path</th>
            <th>Method</th>
            <th>Response</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>/ (root)</td>
            <td>GET</td>
            <td>List of objects</td>
        </tr>
    </tbody>
</table>