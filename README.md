# mysql-hash

Generates a mysql (4.1+) compatible hash.

## Installation

Download the [latest binary release](https://github.com/nutmegdevelopment/mysql-hash/releases/latest)

```bash
wget https://github.com/nutmegdevelopment/mysql-hash/releases/download/1.0.1/mysql-hash_darwin-amd64
mv mysql-hash_darwin-amd64 /usr/local/bin/mysql-hash
chmod u+x /usr/local/bin/mysql-hash
```

## Usage

With an existing password:

```bash
mysql-hash <password>
```

To create a new password:

```bash
mysql-hash -generate
```
