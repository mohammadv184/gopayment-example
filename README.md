# GoPayment Example

This repository includes examples of PCI-compliant UI integrations for online payments with GoPayment Package. Within this demo app, you'll find a simplified version of an e-commerce website

## Supported Integrations

**Golang + Gin Gonic** demos of the following client-side integrations are currently available in this repository:

- [Drivers](https://github.com/mohammadv184/gopayment#list-of-available-drivers)
  - IDPay
  - ZarinPal
  - PayPing
  

## Requirements

Golang 1.17+

## Installation

1. Clone this repo:

```
git clone https://github.com/mohammadv184/gopayment-example.git
```

## Usage

1. Complete `./.env` file with your Gateway Config - Remember to add `http://localhost:3000` as an origin for client key, and merchant account name (all credentials are in string format):

```
{your_driver}_API_KEY="your_API_key_here"
{your_driver}_MERCHANT_ACCOUNT="your_merchant_account_here"
{your_driver}_CLIENT_KEY="your_client_key_here"
```

2. Start the server:

```
go run -v .
```

3. Visit [http://localhost:3000/](http://localhost:3000/) to select an Payment Gateway and complete the payment.

## License

MIT license. For more information, see the **LICENSE** file in the root directory.
