# Project structure based on DDD and clean architucture
 
This is a tested and robust Implementation of clean architecture template. DDD principles can be implied to it easily. Also if you define interfaces for each service in each domain ( these interfaces are called contracts ) you will have hexagonal architecture too. To explain it more, when you implement your services and define interfaces for them in each layer, you are encapsulating them, so you will have a robust app that can be scaled with no worries. 

You should make a folder for each domain with exsisting template in internal folder and put all your external packages in pkg folder; so if you change an external package, the only layers that should be changed are "pkg" and "infrastructure".



### Can be used in microservices or monolith apps. 

