# GopherSentinel

### Objetivos do projeto
O objetivo do projeto é bem simples, ser um bot para discord com poucos comandos para filtrar mensagens que podem ser inapropriadas para crianças. O bot ainda está em desenvolvimento e foi criado com o intuito de aprender Go, uma linguagem na qual sempre tive curiosidade em conhecer.  

### Atenção:
Para utilizar o mesmo você precisa de um arquivo chamado credentials.json na pasta raiz do GopherSentinel.go, o arquivo precisa ter as seguintes chaves:

```json
{
    "BOT_TOKEN": "XXXXXX",
	"APP_ID": "XXXXXX",
	"PUBLIC_KEY": "XXXXXX"
}
```

Caso queira mudar o local que o credentials.json se encontre, passar como parametro para função "ReadCredentialsFile()"