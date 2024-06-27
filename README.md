# PGTest: Your PostgreSQL Connection Wizard 🧙‍♂️✨


PGTest is a lightning-fast, user-friendly CLI tool designed to make PostgreSQL connection testing a breeze for DevOps engineers and database administrators. Say goodbye to connection headaches and hello to instant database diagnostics! 🚀

## 🌟 Features

- **Quick Connection Tests**: Verify your PostgreSQL connections in seconds.
- **Interactive Mode**: Guided input for those who prefer a step-by-step approach.
- **Config File Support**: Store your connection details securely in a YAML file.
- **Secure Password Handling**: Built-in encryption for storing sensitive information.
- **Detailed Diagnostics**: Get helpful troubleshooting tips when connections fail.
- **Cross-Platform**: Works on Windows, macOS, and Linux.

## 🚀 Quick Start

1. Install PGTest:
   ```
   go get github.com/rammanokar/pgtest
   ```

2. Run it:
   ```
   pgtest
   ```

3. Follow the interactive prompts or use command-line flags for quick tests!

## 🛠 Usage

### Command-line flags:

```
pgtest --host localhost --port 5432 --user myuser --dbname mydb
```

### Using a config file:

Create a `config.yaml` file:

```yaml
host: localhost
port: 5432
user: myuser
password: myencryptedpassword
dbname: mydb
sslmode: disable
```

Then run:

```
pgtest --config /path/to/config.yaml
```

## 🔒 Security

PGTest takes security seriously. Passwords are encrypted before being stored in the config file. The encryption key is generated at runtime and never stored.

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for more details.

## 📜 License

PGTest is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

## 🙋‍♀️ Support

Having troubles? Check out our [FAQ](FAQ.md) or open an issue!

---

PGTest: Connect with confidence, test with ease! 💪🔍