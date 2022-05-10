# Notional guide to Hermes

Hermes is a long-running process that reads all channels by default over each node's websocket.  

Channels can and should be filtered, so that hermes can start more quickly. 

## Omni-chain Hermes Quickstart
In order for this to work effectively, you're going to need an archive node for every member of the IBC Gang per [Map Of Zones](https://mapofzones.com). You will also need funds on all relevant chains.  

```bash
mkdir ~/.hermes
cp config.toml ~/.hermes
git clone https://github.com/informalsystems/ibc-rs
cd ibc-rs/relayer-cli
cargo install --path .
```

**In config.toml, make sure that you comment out any chain that you do not wish to relay, or "deny" its channels.**


```bash
hermes keys restore cosmoshub-4 -m "12 or 24 magic words"
hermes keys restore emoney-3 -m "12 or 24 magic words"
hermes keys restore akashnet-2 -m "12 or 24 magic words"
hermes keys restore regen-1 -m "12 or 24 magic words"
hermes keys restore osmosis-1 -m "12 or 24 magic words"
hermes keys restore microtick-1 -m "12 or 24 magic words"
hermes keys restore bitcanna-1 -m "12 or 24 magic words"
hermes keys restore sentinelhub-2 -m "12 or 24 magic words"
hermes keys restore irishub-1 -m "12 or 24 magic words"
hermes keys restore iov-mainnet-ibc -m" 12 or 24 magic words" -p "m/44'/234'/0'/0/0"
hermes keys restore crypto-org-chain-mainnet-1 -m "12 or 24 magic words" -p "m/44'/394'/0'/0/0"
hermes keys restore core-1 -m "12 or 24 magic words" -p "m/44'/750'/0'/0/0"
hermes keys restore columbus-5 -m "12 or 24 magic words" -p "m/44'/330'/0'/0/0"
hermes keys restore kava-9 -m "12 or 24 magic words" -p "m/44'/459'/0'/0/0"
hermes keys restore impacthub-3 -m "12 or 24 magic words"
``` 




### Snippets
```bash
hermes query packet commitments osmosis-1 transfer channel-0
hermes tx raw packet-recv cosmoshub-4 osmosis-1 transfer channel-0
```

### Tips
When congestion accidentally happens, separating the stuck channels from the running hermes will prevent it from harming other networks' relaying. We suggest commenting it out from the config file and running it on go-relayer/hermes or setting the 'hermes clear packets' commands in a loop.
p/s: Delay/Performance impact when having stuck channels in config, Hermes v0.14.1
