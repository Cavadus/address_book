# Address Book

Address Book is a RESTful API utilizing a MySQL database and CSV import/export capabilities.

# Installation

```
git clone https://github.com/Cavadus/address_book.git
cd address_book
```

At this point edit line #8 in main.go for your MySQL server connection details (username, password, database name).

```
go install
```

## Dependencies

This application uses the following dependencies (Gopkg files will ensure install when running "go install"):

```
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
```

## Database
The address book contains four VARCHAR(255) fields in addition to an auto-incrementing primary key ID and obligatory timestamp.  These four columns are:
1. First name (fname)
2. Last name (lname)
3. E-mail address (email)
4. Phone number (phone)
```
+--------------+-------------------+-------+------+--------------------+----------------+--+
|    Field     |       Type        | Null  | Key  |      Default       |     Extra      |  |
+--------------+-------------------+-------+------+--------------------+----------------+--+
| id           | int(10) unsigned  | NO    | PRI  | NULL               | auto_increment |  |
| fname        | varchar(255)      | NO    |      | NULL               |                |  |
| lname        | varchar(255)      | NO    |      | NULL               |                |  |
| email        | varchar(255)      | YES   |      | NULL               |                |  |
| phone        | varchar(255)      | YES   |      | NULL               |                |  |
| createdDate  | datetime          | YES   |      | CURRENT_TIMESTAMP  |                |  |
+--------------+-------------------+-------+------+--------------------+----------------+--+
```
## SQL Statement
If you wish:
```
CREATE DATABASE pizza_hut;
USE pizza_hut;
CREATE TABLE address_book (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fname VARCHAR(255) NOT NULL,
    lname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NULL,
    phone VHARCHAR(255) NULL
);
```

#	



```python
import foobar

foobar.pluralize('word') # returns 'words'
foobar.pluralize('goose') # returns 'geese'
foobar.singularize('phenomena') # returns 'phenomenon'
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
