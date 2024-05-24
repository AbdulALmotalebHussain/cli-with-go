-- Create a new database called "mydatabase"
CREATE DATABASE mydatabase;

-- Use the new database
USE mydatabase;

-- Create a new table called "mytable"
CREATE TABLE mytable (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  age INT NOT NULL,
  email VARCHAR(100) NOT NULL,
  PRIMARY KEY (id)
);
DESCRIBE mytable;
-- Insert some sample data into the table
INSERT INTO mytable (name, age, email) VALUES
('John Doe', 25, 'john.doe@email.com'),
('Jane Smith', 30, 'jane.smith@email.com'),
('Bob Johnson', 40, 'bob.johnson@email.com');
SELECT * FROM mytable;






DROP DATABASE costmer;
CREATE DATABASE costmer;
CREATE TABLE employee(
empid int,
 name VARCHAR(25)NOT NULL,
lastname VARCHAR(25)NOT NULL ,  
 salary decimal(5,3),
    dep VARCHAR(50)DEFAULT 'SE',
    CONSTRAINT employee_PK PRIMARY KEY(empid),  
    CONSTRAINT employee_CHECK CHECK (salary>100),
);
    DESCRIBE employee;
    INSERT INTO  employee(empid,name,lastname,salary,dep) VALUES
    (1,'ABD','HUSSAIN',1000,'MATH');
    SELECT * FROM employee;