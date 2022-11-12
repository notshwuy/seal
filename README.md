<img width="100px" src="https://www.emojiall.com/images/240/microsoft/1f9ad.png" align="left">

### Seal

A tool to run commands with monitoring features.

##### Features

- **`seal run --load-env`** loads `.env` file.
- **`seal run --auto-restart`** restarts the application when goes down.
- **`seal run --output-logs`** outputs log files. (*in development*)

##### Usage

> By now, [Go >= 1.17](https://go.dev/dl/) is required to install and use this tool. 

```dockerfile
# Install Seal CLI
RUN go install github.com/sxhk0/seal/cmd/seal@latest

# Run your script
RUN seal run -- node src/index.js

# Run your script with "Auto Restart" and "Load .env"
RUN seal run -r -e -- node src/index.js
#                  ^^ your command execution should be after --
```

##### License

[MIT](/LICENSE) &copy; Itallo Gabriel
