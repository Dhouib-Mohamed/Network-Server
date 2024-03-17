# Protohackers Challenges in Go

This repository contains solutions for some of the challenges in Protohackers, a casual programming challenge where you create servers for network protocols. Protohackers is created by James Stanley.

## About Protohackers

Protohackers is a casual programming challenge in which you create servers for network protocols.

- You are given a network protocol spec.
- You implement the server and input its IP address and port number.
- Your server is automatically checked.
- Once all checks pass, your solve time goes onto a global leaderboard.

The challenges start easy and get harder over time.

For more information, visit the [Protohackers website](https://protohackers.com/).

## Solved Problems

The following problems have been solved and their solutions are included in this repository:

1. **Smoke Test**
2. **Prime Time**
3. **Means To an End**

## How to Use

Each challenge is implemented in a separate directory inside src/main. 
To run the server you can just run the main.go file:

```bash
go run network-challenges
```

And it will run all the solved challenges concurrently .
To customize the ports where each challenge will run, you can consult .example.env and create your own env customized file.

To run a challenge, follow the instructions in the respective directory.

## Contributing

Contributions are welcome! If you have a solution for a challenge that is not yet included in this repository, feel free to submit a pull request.

## Contact

For any inquiries about Protohackers, you can email James Stanley at jes@protohackers.com, join the Discord group, or follow @Protohackerscom on Twitter.
