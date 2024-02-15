<!--
parent:
  order: false
-->

<div align="center">
  <h1> Ethermint </h1>
</div>

[Entangle](https://www.entangle.fi/) blockchain node based on Ethermint - a scalable and interoperable Ethereum library, built on Proof-of-Stake with fast-finality using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

**Note**: Requires [Go 1.19+](https://golang.org/dl/)

## Become a Validator

You can check the instructions to [Run a Node](https://entangle-protocol.gitbook.io/welcome/getting-started/infrastructure-guides/run-a-validator-node).

### tl;dr

#### Hardware Requirements

|         | Mainnet                      | Testnet                      |
|---------| -----------                  | -----------                  |
| CPU     | 4 or more physical CPU cores | 2 or more physical CPU cores |
| RAM     | 16GB+                        | 8GB+                         |
| Storage | 500GB+ SSD                   | 250GB+ SSD                   |
| Network | 100mbps+                     | 10mbps+                      |
| OS      | Ubuntu 22.04                 | Ubuntu 22.04 or macOS        |

#### Installation and run

1. Install Git, Golang (minimum version is 1.19), make, jq, Python, golangci-lint, and Solc-JS

2. Clone the repository, or check out the latest [release](https://github.com/Entangle-Protocol/entangle-blockchain/releases):

```bash
git clone https://github.com/Entangle-Protocol/entangle-blockchain
cd entangle-blockchain
```

3. Install the application:

```bash
make install
```

4. <p id="item4">Run the account generation script with your key name and encryption password as the parameters below:</p>

```bash
sh init_key.sh <password>
```

5. Execute the script for downloading and unpacking data from previous blocks:

```bash
sh get_snapshot.sh
```

6. Initiating the full blockchain node will commence the process of synchronizing the current block state:

```bash
sh run_node.sh
```

7. Receiving NGL Tokens to the Account

    To add an account to MetaMask, you need to obtain the private key of your account. You can do this using the command:

    ```bash
    entangled keys unsafe-export-eth-key <key_name>
    ```

    The obtained private key can be used to import the account into MetaMask. Afterward, you can acquire NGL tokens for the account using the Entangle faucet via the [Discord](https://discord.com/invite/entanglefi).

8. Starting a Validator
    
    As the `amount`, specify the number of tokens you are staking for the Validator, and as the `from`, use the key name used in the [Validator Account Generation step](#item4).

    ```bash
    entangled tx staking create-validator \
    --amount="5000000000000000000aNGL" \
    --pubkey=$(entangled tendermint show-validator) \
    --moniker="validator" \
    --chain-id=entangle_33133-1 \
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1" \
    --gas=500000 \
    --gas-prices="10aNGL" \
    --from=<key_name>
    ```

## Community

- [Official Website](https://www.entangle.fi/)
- [Twitter](https://twitter.com/Entanglefi)
- [Discord](https://discord.com/invite/entangle)

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/Entangle-Protocol/entangle-blockchain/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.
