### Deploy database

Open a query window in SQL Management Studio and run the sql script in this folder.

You will need to update the user "test" password.

Update the connection string on line 22 in file `server/middleware/middleware.go`

###WARNING

Note: Storing password in code is not ideal and it is not recommended for production environments. This is done for for testing purpose and not complicate the code. 
 