# Golang Aggregator Service

This application currently will pull all MLB transactions over the past week and store them into a PostgreSQL database

## API

`http://lookup-service-prod.mlb.com/json/named.transaction_all.bam?sport_code='mlb'&start_date='<YYYYMMDD>'&end_date='<YYYYMMDD>'`


**example**

http://lookup-service-prod.mlb.com/json/named.transaction_all.bam?sport_code='mlb'&start_date='20190414'&end_date='20190419'
