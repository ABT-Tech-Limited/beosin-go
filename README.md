# Beosin Go SDK

Beosin API 的官方 Go SDK，用于区块链地址和交易的风险评估。

## Installation

```bash
go get github.com/ABT-Tech-Limited/beosin-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    beosin "github.com/ABT-Tech-Limited/beosin-go"
)

func main() {
    // Create client
    client := beosin.NewClient("your-app-id", "your-app-secret")
    ctx := context.Background()

    // Query account balance
    balance, err := client.GetAccountBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance: %d credits\n", balance.Data.SurplusIntegral)

    // Address risk assessment
    resp, err := client.EOAAddressRiskAssessment(ctx, &beosin.AddressRiskRequest{
        ChainID: beosin.ChainETH,
        Address: "0x...",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Risk Level: %s\n", resp.Data.IncomingLevel)
}
```

## Features

| Method | Description |
|--------|-------------|
| `GetAccountBalance` | Query account credits |
| `EOAAddressRiskAssessment` | EOA address risk assessment |
| `DepositTransactionAssessment` | Deposit transaction risk assessment |
| `WithdrawalTransactionAssessment` | Withdrawal transaction risk assessment |
| `MaliciousAddressQuery` | Query malicious address tags |
| `VASPQuery` | Query VASP information |
| `BlackAddressScreening` | Black address screening |

V4 API:

| Method | Description |
|--------|-------------|
| `V4EOAAddressRiskAssessment` | V4 EOA address risk assessment |
| `V4DepositTransactionAssessment` | V4 deposit transaction assessment |
| `V4WithdrawalTransactionAssessment` | V4 withdrawal transaction assessment |

## Options

```go
client := beosin.NewClient(appID, appSecret,
    beosin.WithTimeout(60 * time.Second),
    beosin.WithDebug(true),
    beosin.WithBaseURL("https://custom-api.example.com"),
)
```

## Supported Chains

`ChainETH`, `ChainBSC`, `ChainPolygon`, `ChainArbitrum`, `ChainOptimism`, `ChainAvalanche`, `ChainTron`, `ChainSolana`, `ChainBTC`, `ChainTON`, `ChainAptos` and more.

See [types.go](types.go) for the full list.

## License

MIT
