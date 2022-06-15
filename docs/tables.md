# Tables

This guide contain infomation of tables used in `ProjectService`.

## Table List
```mysql
mysql> show tables;
+------------------+
| Tables_in_raffle |
+------------------+
| project          |
| user             |
| user_project     |
| user_wallet      |
+------------------+
```

### Table1. project
```mysql
mysql> show columns from project;
+-----------------+--------------+------+-----+---------+----------------+
| Field           | Type         | Null | Key | Default | Extra          |
+-----------------+--------------+------+-----+---------+----------------+
| id              | int          | NO   | PRI | NULL    | auto_increment |
| name            | varchar(50)  | NO   |     | NULL    |                |
| chain_id        | int          | NO   |     | NULL    |                |
| raffle_contract | varchar(100) | NO   |     | NULL    |                |
+-----------------+--------------+------+-----+---------+----------------+
```