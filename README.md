Prerequisites -
You would need VS code, go programming language, postman and Heidi SQL installed in your system

Database Structure -
We are using mySQL database and the table structure is as below -
Database name - cards, table names - Accounts, Transactions and OperationTypes

To connect to the database -
Network type - Maria DB or MySQL (TCP/IP)
Library - libmariadb.dll
Hostname/IP - 127.0.0.1
User - gurleen
Password - Jaibir2021123#
Port - 3306

After cloning the code in a folder in your machine - 
Open command prompt and navigate to the source code folder and run the "go run myserver.go" command; that should start our server
Open the postman app and we are ready to test our code!

The functions we have in the code - 
getAccounts function - retrieves all rows in the Accounts table (Get method) 
getAccount function - retrieves a single account based on the account_id (Get method)
createAccount function - creates a new record in the Accounts table (Post method)
getTran function - retrieves all the rows in the Transactions table (Get method)
createTran function - creates a new record in the Transactions table (Post method)
