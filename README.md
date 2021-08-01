# golang-basic-skeleton

Basic skeleton for golang projects

1. You can place the README texts here, like setup steps of project, or any other information about the project that you think would be helpful for other developers/users.

2. Create `.env` file, where we will keep all environment related constants, like database username, password, secret key of JWT if any, etc...

3. Create the "main.go" file in root, this will server as the project/web-serve entry point.

4. Create a "Makefile" this will contain some basic commands to run the project.

5. Create these folders:    
    1. Controllers: contains controller that accept a request and call a particular service to process it.
    
    2. Services: contains service that has the logic for doing all the procces like manipulating DB and giving back the desired result.
    
    3. Models: contains the struct for storing data in DB or fetching data from DB.
    
    4. DB: keep your DB related operation here, these can used in services to fetch/update/insert/delete data from DB.
        *** Place DB connection file as well inside this folder.

    5. Routes: keep all your routes here and pass the name of controllers.

    6. Config: keep all your configuration things here like, fetching variables from `.env` file, or even DB config can be placed here.

    7. Constants: keep all constants here, can also categorise them in separate files.

    8. Utils: keep commonly used functions here like capitalizing first letter, etc...

So, here we are ready with one of the best basic folder structure for a web-server development in golang there are many way.