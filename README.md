# Cipher CLI in Go

This project is a Command Line Interface (CLI) tool written in Go for generating and managing ciphers. The tool allows users to cipher and decipher messages with varying levels of complexity. Cipher keys are stored temporarily and expire after a specified duration. The project includes encryption, user authentication, and Dockerization.

## Features

- Generate ciphers with customizable complexity levels.
- Encrypt and store cipher keys securely for a limited time (10 minutes).
- Cipher and decipher messages using stored keys.
- User authentication for secure access.
- Dockerized for easy deployment.
- Configurable via a YAML configuration file.
