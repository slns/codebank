### Rodar servidor Go
> go run main.go

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

### Saber o ip do container
> docker inspect nome_do_container -f "{{json .NetworkSettings.Networks }}"

### Base de dados
> [Link BD](http://localhost:9000)

#### User/Pass
```
admin@user.com

 123456
 ```
#### gRPC in Mode Cliente(Ir ao bash do projeto codebank_aluno)
```
evans -r repl -p=50052
```

##### Chamar dentro do Evans o Payment
```
call Payment
```