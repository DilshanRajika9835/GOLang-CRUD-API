Drop DATABASE  If EXISTS  GOLand;
CREATE DATABASE GOLand;
USE GOLand;
CREATE TABLE Customer
(
    ID      VARCHAR(11)    NOT NULL,
    Name    VARCHAR(30)    NOT NULL,
    Address VARCHAR(30)    NOT NULL,
    Salary  decimal(10, 2) not null,
    CONSTRAINT PRIMARY KEY (ID)
);
