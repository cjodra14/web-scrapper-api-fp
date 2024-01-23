# web-scrapper-api-fp


RUN FLAGD

```
docker run   --rm -it   --name flagd   -p 8013:8013   -v $(pwd):/etc/flagd   ghcr.io/open-feature/flagd:latest start   --u
ri file:/etc/flagd/conf.flagd.json
```


```
docker build -t web-scrapper .

docker run web-scrapper
```

Post scrapping Uri example

```
http://localhost:8080/scrap?uri=https://www.foroparalelo.com/deportes/carlos-sainz-ganador-del-dakar-2024-a-1199356

```


