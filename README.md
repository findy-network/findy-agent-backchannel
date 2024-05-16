# findy-agent-backchannel

[![test](https://github.com/findy-network/findy-agent-backchannel/actions/workflows/test.yml/badge.svg?branch=dev)](https://github.com/findy-network/findy-agent-backchannel/actions/workflows/test.yml)

> Findy Agency is an open-source project for a decentralized identity agency.
> OP Lab developed it from 2019 to 2024. The project is no longer maintained,
> but the work will continue with new goals and a new mission.
> Follow [the blog](https://findy-network.github.io/blog/) for updates.

Findy Agency backchannel for testing agency interoperability using Aries Agent Test Harness.

Running the tests:

```
cd env
make aath-up
```

The backchannel server sources are generated with [openapi-generator](https://openapi-generator.tech/) using [AATH openapi specification](https://github.com/hyperledger/aries-agent-test-harness/blob/main/docs/assets/openapi-spec.yml).
