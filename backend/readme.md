#Backend

####database
set it up on ubuntu 14.04

installed mongo from [this link](https://docs.mongodb.org/manual/tutorial/install-mongodb-on-ubuntu/)

thought about using mongo but our data is probably going to be pretty structured.
I am not sure that we will need to add a lot of unexpected data. NoSQL (or at least
mongo) seems to shine when there is unstructured or diverse data being added. I 
think that our application will likely have pretty structured data.

I ended up using sqlite because it is fast, light and easy to use

installed sqlite with apt-get sqlite3 libsqlite3-dev
