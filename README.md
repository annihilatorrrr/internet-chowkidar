# Internet Chowkidar
A tool that aggregates data on Internet Service Provider (ISP)
blockages. This initiative is designed to monitor and record instances
where certain websites are blocked or restricted by different ISPs
across various regions.

Read more at this [page](https://gnulinuxindia.sh/blog/plans-for-june-24/).

## Prerequisites
- Go 1.22
- wire
- sqlc

## Getting Started
1. Clone the repository
```bash
git clone https://github.com/gnulinuxindia/internet-chowkidar.git
```

2. Run code generation
```bash
sqlc generate
go generate ./...
```

3. Run the application
```bash
go run main.go
```
