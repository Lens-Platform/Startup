## Create Service Local Database
---

From this directory run the below operation
```bash
docker-compose up
```

Navigate to `localhost:16543` and enter the following sets of credentials
```bash
email: present in .env file for pgadmin
password: present in .env file for pgadmin
```

When logged in, on the right side of the pane, right click on the servers options and create a server.
```bash
1. Add server name ... call it the same as the db name in .env file
2. Click on connections tab
    A. go on local terminal and run ifconfig | grep int
    B. copy the address that follows the following semantic
       1A. inet 10.101.1.134 netmask 0xffffff00 broadcast 10.101.1.255
       1B. so copy 10.101.1.134 from the above example into the host name/connection
           field
    C. select maintenance database to be databse specified in .env file
    D. username and password are same as those present in .env file
    E. click save
3. Click on servers --> test --> schemas --> right click Tables --> View/Edit Data --> All Rows
```
---

NOTE: please ensure to reference the .env file and configure to your liking. Additionally, the docker_postgres_init.sql is for example use. For
 your service, define your own schema and seed your local database based on your service needs.
