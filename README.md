
# mysql-hash

Generates a mysql (4.1+) compatible hash.

### Installation

Download the [latest binary release](https://github.com/nutmegdevelopment/mysql-hash/releases/latest)

Unzip it and add it to a folder in your path.

```
wget https://github.com/nutmegdevelopment/mysql-hash/files/591789/mysql-hash_darwin-amd64_v1.0.0.zip
unzip mysql-hash_darwin-amd64_v1.0.0.zip
mv mysql-hash /usr/local/bin/mysql-hash
chmod u+x /usr/local/bin/mysql-hash
```

### Usage

With an existing password:
```
mysql-hash <password>
```

To create a new password:
```
mysql-hash -generate
```
