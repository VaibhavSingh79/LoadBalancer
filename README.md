<h1>Go-Based Load Balancer 🚀</h1>

Welcome to the Go-Based Load Balancer project,a lightweight yet powerful HTTP load balancer written from scratch in Go! 
This project is designed to efficiently distribute incoming traffic across multiple backend servers, ensuring high availability, fault tolerance, and balanced workload distribution. 🌐✨


<h2>Features ✨</h2>
<ul>
<li>🌍 Round Robin Load Balancing: Implements the Round Robin algorithm to evenly distribute HTTP requests across backend servers.
<li>🔗 Reverse Proxy: Leverages Go's httputil.ReverseProxy for seamless request forwarding to backend servers.
<li>⚙ Fault Tolerance: Excludes servers that are not alive, ensuring requests are only sent to healthy servers.
<li>🚀 Concurrent Support: Fully supports Go's goroutines for handling multiple simultaneous requests with optimal performance.
<li>🪶 Lightweight and Simple: A minimal Go-based application that's easy to extend and deploy for modern distributed systems.
</ul>


<h2>How It Works 🤔</h2>

<h3>🌱 Initialization:</h3>
<ul>
<li>A pool of backend servers is defined, each represented by a SimpleServer instance.
The load balancer, running on a specified port, listens for client requests.</li>

<h3>🔀 Reverse Proxy:</h3>
<li>The incoming request is forwarded to a backend server using ReverseProxy. The selected server processes the request and sends the response back to the client seamlessly.</li>

  
<h3>📊 Round Robin Algorithm:</h3>

<li>The load balancer utilizes a thread-safe Round Robin algorithm to achieve fair request distribution across all available backend servers.</li>

<h3>🚧 Fault Handling:</h3>

<li>Servers are conditionally excluded during the selection loop if they are marked as unhealthy. This ensures high availability by redirecting traffic to only active servers.</li>
</ul>
