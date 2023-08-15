### hacktiv8 workshop

#### readme first
- ```$ mv .env.sample .env```
- setting .env in your local

#### run service order
- ``` $ go run services/orderservice/cmd/server/main.go```

#### service order via docker
- ```make run-order-service```

#### run service payment
- ``` $ go run services/orderservice/cmd/server/main.go```

#### service payment via docker
- ```make run-payment-service```

## API SPEC
#### Creating order
POST /v1/order
request
```{
    "price": 2000,
    "items": "samsung",
    "customer_id": 100
}
```
response
```
{
    "id": 1
}
```
#### Get list order
GET /v1/order
response
```
[
    {
        "id": 1,
        "customer_id": 1,
        "items", "iphone 14",
        price: 5000,
        payment_status: "UNPAID"
    }
]
```


#### Make payment
POST /v1/payment
request
```
{
    "order_id": 1
}
```
response
```
{
    "id": 3
}
```

#### Get List payment
GET /v1/payment
```
[
    {
        "id": 1,
        "order_id": 7,
        "payment_status": "PAID"
    },
    {
        "id": 2,
        "order_id": 1,
        "payment_status": "PAID"
    },
    {
        "id": 3,
        "order_id": 8,
        "payment_status": "PAID"
    }
]
```

