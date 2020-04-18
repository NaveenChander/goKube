# goKube


Docker Build Command
------ ----- -------
docker build -t gokube .



Docker Run Command
------ --- -------
[docker run -p 8000:8000 gokube] -> This Docker container doesnt have a name. 
[docker run -p 8000:8000 --name MyNewDocker gokube] -> This Docker has a name [MyNewDocker]. You can use this name to stop or remove a containerg  

[docker image ls]  -> List all docker images
[docker ps -a] -> List all Docker containers with name and port
[docker stop [containerName]] 


MakeFile Commands
-------- --------
Make Build
Make Run