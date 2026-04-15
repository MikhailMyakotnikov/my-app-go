CREATE USER IF NOT EXISTS 'mikhail'@'%' IDENTIFIED BY '${userPassword}';

GRANT ALL PRIVILEGES ON test_db.* TO 'mikhail'@'%';
GRANT ALL PRIVILEGES ON prod_db.* TO 'mikhail'@'%';

FLUSH PRIVILEGES;