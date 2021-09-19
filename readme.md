### Levantar os conteineres
> docker-compose up -d

### Ver conteineres levantados
> docker-compose ps

### Entrar num container
> docker exec -it nome_do_container bash

### Iniciar progeto golang, pacote onde vou guadar
> go mod init githib.com/slns/codebanck

### Download imports
> go get -u nome_do_pacote


docker inspect 5912ceda5a23 -f "{{json .NetworkSettings.Networks }}"