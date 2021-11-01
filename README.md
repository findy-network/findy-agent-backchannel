# findy-agent-backchannel

[![test](https://github.com/findy-network/findy-agent-backchannel/actions/workflows/test.yml/badge.svg?branch=dev)](https://github.com/findy-network/findy-agent-backchannel/actions/workflows/test.yml)

Findy Agency backchannel for testing agency interoperability using Aries Agent Test Harness.

Running the tests:

```
cd env
make aath-up
```

The backchannel server sources are generated with [openapi-generator](https://openapi-generator.tech/) using [AATH openapi specification](https://github.com/hyperledger/aries-agent-test-harness/blob/main/docs/assets/openapi-spec.yml).
