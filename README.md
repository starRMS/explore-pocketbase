# Exploring PocketBase

> [!NOTE]<br>
> In active development

# Table of Contents
- [Introduction](#introduction)
- [Components](#components)
  - [Observability](#1-observability)

## Introduction
This project serves as an exploration to [pocketbase](https://pocketbase.io) and template for using pocketbase integrated with other extra components.

## Components

### 1. Observability
We want our application to be observable. Hence we integrated [opentelemetry](https://opentelemetry.io/) into the application with the help of [Jaeger](https://www.jaegertracing.io/).
##### How to start - Jaeger
To start Jaeger please refer to this [link](https://www.jaegertracing.io/docs/1.52/getting-started/) to start jaeger suiting your usecase. For starter, we recommend starting jaeger with its all-in-one mode that is simply running the command below:
```bash
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.52
```

Then you can head to [http://localhost:16686](http://localhost:16686) to see the jaeger UI.

