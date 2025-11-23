# note commands

```bash
# make the swarm
docker swarm init --listen-addr 0.0.0.0:2377
docker stack deploy -c docker-compose.yml orderApp 
docker swarm join --token <tocken> 192.168.1.202:2377   

# check the functionality of the services
docker stack ls
docker stack ps orderApp
docker service logs <serviceid>

# end the swarm
docker swarm leave
docker swarm leave --force
```
