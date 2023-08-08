# GopherSentinel

### Objetivos do projeto
O objetivo do projeto é bem simples, ser um bot para discord que serve como um "vigía" para filtrar mensagens, imagens e arquivos que podem ser inapropriadas para crianças, ou para o publico em geral. O bot ainda está em desenvolvimento e foi criado com o intuito de aprender e praticar Go, uma linguagem na qual sempre tive curiosidade em conhecer.  

### Atenção:
Para utilizar o mesmo você precisa de um arquivo chamado credentials.json na pasta raiz do GopherSentinel.go, o arquivo precisa ter as seguintes chaves:

```json
{
 "BOT_TOKEN": "XXXXXX", // Token bot discord
 "APP_ID": "XXXXXX", // Bot App id
 "PUBLIC_KEY": "XXXXXX", // Discord public key
 "GOOGLE_USER_PROJ": "XXXXXX", // Google ProjectId
 "GOOGLE_BEARER_TOKEN": "XXXXXX" // Google bearer token (gcloud auth print-access-token)
}
```

Caso queira mudar o local que o credentials.json se encontre, passar como parametro para função "ReadCredentialsFile()"

### Features adicionadas

* Detecção e exclusão de mensagens que contenham palavrões
* Reconhecimento e exclusão de imagens com conteudo, adulto, ofensivo, racista e etc.