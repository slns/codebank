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


docker inspect nome_do_container -f "{{json .NetworkSettings.Networks }}"