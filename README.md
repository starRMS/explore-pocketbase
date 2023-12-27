# Exploring PocketBase

> [!NOTE]<br>
> In active development

# Table of Contents
- [Introduction](#introduction)
- [Components](#components)
- [Getting Started](#getting-started)

## Introduction
This project serves as an exploration to [pocketbase](https://pocketbase.io) and template for using pocketbase integrated with other extra components.

## Components
The components included are:
1. [Pocketbase](https://pocketbase.io) - The main application
2. Observability Stack:
    - [Opentelemetry](https://opentelemetry.io/)
    - [Jaeger](https://www.jaegertracing.io/)
    - [Prometheus](https://prometheus.io/)
    - [Grafana](https://grafana.com/)

## Getting Started
You can start this project by using the docker compose option. It will start all the components including the app and observability stack.
You can simply run:
```bash
docker compose up --build -d
```
In result, the pocketbase app admin page will be available in [http://localhost:8090/_/](http://localhost:8090/_/). Meanwhile, the Jaeger client will be available at [http://localhost:16686](http://localhost:16686), Prometheus will be available at [http://localhost:9090](http://localhost:9090) and the Grafana will be available at [http://localhost:3000](http://localhost:3000).
