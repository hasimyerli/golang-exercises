# Testing lambda functions locally with Sam

### How to work?

Open the terminal.

Step 1
> cd binance-ticker

> go mod init binance-ticker

> go mod tidy


Step 2
> cd btcturk-ticker

> go mod init btcturk-ticker

> go mod tidy



Step 3

> sam build 


Step 4
> sam local start-api


### Output

http://localhost:3000/binance/ticker
Output:
```json
[
    {
        "symbol": "ETHBTC",
        "priceChange": "-0.00026400",
        "priceChangePercent": "-0.351"
    },
    ...
]
```

http://localhost:3000/btcturk/ticker
Output:
```json
{
    "data": [
        {
            "pair": "BTCTRY",
            "daily": -160,
            "dailyPercent": -0.1
        },
        ...
    ]
}
```