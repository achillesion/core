#!/usr/bin/env bash

rm -rf ~/.ticketapp

dtd init a --chain-id testchain
echo "12345678" | dtcli keys add test1
echo "12345678" | dtcli keys add test2

dtd add-genesis-account $(dtcli keys show test1 -a) 10000000000000000000000000stake,1000000legends
dtcli config output json
dtcli config indent true
dtcli config trust-node true

echo "12345678" | dtd gentx --name test1

echo "Collecting genesis txs..."
dtd collect-gentxs

echo "Validating genesis file..."
dtd validate-genesis

dtd start 