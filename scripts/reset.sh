#!/bin/bash 


docker-compose -f docker-compose.dev.yml down -v
docker-compose -f docker-compose.prod.yml down -v
