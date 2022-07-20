# API - GOLANG - Cambio 2

## Operações possíveis

1. Inserir um deposito
2. Listar todos os depositos
3. Listar deposito por id
4. Consultar saldo total
5. Consultar saldo convertido em moeda do tipo:
    - EUR : euro
    - GBP : libra
    - USD : dolar
    - JPY : iene


## Funções do Gin
- *Criar o nosso server*
- *Axilia na criação das rotas, instanciar os controllers, devolver reposta para o usuário, etc.*

## GORM
- *É uma biblioteca ORM feita em Go que possibilita relacionar os modelos descritos da aplicação às tabelas do banco de dados relacional.  Além disto, o GORM permite realizar consultas no banco de dados SQL utilizando os modelos definidos na camada de modelos. Esta camada de ORM simplifica o desenvolvimento e permite que o código possua mais similaridade com a modelagem do banco de dados.*
```fonte: https://bdm.unb.br/bitstream/10483/27278/1/2020_LucasCamposJorge_tcc.pdf, (p.31)```

## Banco de dados escolhido
- *MySQL*

## Observação
- Precisa criar um arquivo .env com esses dados:
``` 
DB_USER= <seu user>
DB_PASS= <sua senha>
DB_HOST=localhost
DB_NAME= <nome database>
 
```
