# GoLedger Challenge - Besu Edition
 
Simple application that will interact with a Besu node to transact in a smart contract, check the value of a smart contract variable and sync that value to an external database.

To accomplish that, we recommend you use a UNIX-like machine (Linux/macOS). Besides that, we will need to install NPM/NPX, Hardhat and Docker.

## Install the prerequisites

- Install NPM and NPX (https://www.npmjs.com/get-npm)
- Install Hardhat (https://hardhat.org/getting-started/)
- Install Docker and Docker Compose (https://www.docker.com/)
- Install Besu (https://besu.hyperledger.org/private-networks/get-started/install/binary-distribution)
- Install Go (https://golang.org/dl/)

## Set up the environment

To set up the environment, you need to clone this repository. Make sure you have installed the requirements. To set up the environment, you need to run the following commands:


1. Clone this repository
2. Enter into folder:
```bash
cd goledger-challeng-besu
```
3. Enter into besu folder:
```bash
cd besu
```
4. Install hardhat:
```bash
npm install --save-dev hardhat
```
5. Start scripts that will setup besu network:
```bash
./startDev.sh
```
6. Install dependecy:
```bash
go mod tidy
```
7. Run server:
```bash
go run main.go
```

**Note**
This server running in port 8080, you can change this in the code if you want.

# The app

- **Gin Framework**: Used for building the REST API.
- **SQLite**: Chosen for its simplicity and efficiency as the database.
- **GORM**: ORM for interacting with the database.

## Endpoints:

- **Set**: Interacts with the Besu network to insert a value into a smart contract.
- **Get**: Interacts with the Besu network to retrieve the value stored in the smart contract.
- **Sync**: Interacts with the Besu network to retrieve the value from the smart contract and insert/update it in the database.
- **Check**: Interacts with the Besu network to retrieve the value from the smart contract, compares it with the value stored in the database, and performs a validation.

You can check these endpoints here: .

## Application Structure:

- **besu**: Handles interactions with the Besu network.
- **config**: Manages the app's configurations, including database and logger initialization.
- **db**: Contains the SQLite database file.
- **handler**: Manages HTTP request handlers.
- **router**: Initializes the API routes.
- **schemas**: Defines the application entities.
