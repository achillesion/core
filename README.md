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
emcli keys add bob

# Copy the `Address` output here and save it for later use
emcli keys add alice

# Add both accounts, with coins to the genesis file
emd add-genesis-account $(emcli keys show bob -a) 1000nametoken,1000jackcoin
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

- write tests for modules
- percentage markup with no floats
- testing testnets
- swagger file for restapi
- keyserver adaptation for jwt Token and
