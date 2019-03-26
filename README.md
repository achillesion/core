# Tic_Mar

Tic mar hehe

## Instal into Bin

```
make install
```

## Run the chain

```
# Initialize configuration files and genesis file
emd init --chain-id testchain

# Copy the `Address` output here and save it for later use
emcli keys add jack

# Copy the `Address` output here and save it for later use
emcli keys add alice

# Add both accounts, with coins to the genesis file
emd add-genesis-account $(emcli keys show jack -a) 1000nametoken,1000jackcoin
emd add-genesis-account $(emcli keys show alice -a) 1000nametoken,1000alicecoin

# Configure your CLI to eliminate need for chain-id flag
emcli config chain-id testchain
emcli config output json
emcli config indent true
emcli config trust-node true
```

## Lint

To lint the files

```
make get-linter
make lint
```

## Things to do:

- Makefile
  - build & install
- market module coding
- write tests
