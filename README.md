# GoLoadBalancer

A simple, efficient HTTP Load Balancer written in Go using reverse proxies.
This project demonstrates round-robin load balancing across multiple backend servers, with intelligent MIME handling for JavaScript modules.

---

## Features

* 🔁 **Round-Robin Load Balancing**: Distributes requests evenly across all healthy servers.
* ✅ **Health Check**: Skips servers that are unresponsive or return error status codes.
* 🔁 **Reverse Proxy**: Forwards requests to actual backend servers seamlessly.
* 📦 **MIME Fixes**: Ensures `.js` files are served with the correct `Content-Type`.
* 💡 **Concurrency-Safe**: Uses mutexes to manage round-robin state.

---

## 📦 Getting Started

### 🔧 Prerequisites

* Go 1.18 or higher installed

### 🛠️ Installation

```bash
git clone https://github.com/ayuuuuu0-0/loadbalancer_go.git
cd loadbalancer_go
go run main.go
```

### 🖥️ Example

Update your `main.go` with the target backend servers:

```go
servers := []Server{
	newSimpleServer("https://www.instagram.com"),
	newSimpleServer("https://www.facebook.com"),
	newSimpleServer("https://college-cart.vercel.app"),
	newSimpleServer("https://ayushranjan.tech"),
}
```

Run the load balancer:

```bash
go run main.go
```

Visit [http://localhost:8000](http://localhost:8000) in your browser. Your requests will be forwarded to the backend servers in round-robin order.

---

## 📁 Project Structure

```bash
.
├── main.go        # Load balancer core logic
└── README.md      # Project documentation
```

---

## 🛡️ Health Check Mechanism

Each server is periodically checked for availability via a GET request. If a server responds with an error (status ≥ 400) or times out, it is skipped in the round-robin cycle.

---

## 🧪 Test It Locally

You can test it with mock servers like:

```bash
python3 -m http.server 9001
python3 -m http.server 9002
```

Then point the load balancer to:

```go
newSimpleServer("http://localhost:9001"),
newSimpleServer("http://localhost:9002"),
```

---


## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

---

## 📄 License

[MIT License](LICENSE)

---

## 👨‍💻 Author

Made by Ayush
---

