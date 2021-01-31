drop database golangSimpleApplication;
create database golangSimpleApplication;
use golangSimpleApplication;




CREATE TABLE customer(
	id VARCHAR(225) NOT NULL,
	name VARCHAR(225) NOT NULL,
	address VARCHAR(225) NOT NULL,
	nic VARCHAR(15) NOT NULL,
	contact int(10),
	CONSTRAINT PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;