# 👺 TODOOM API 👺 WIP

## 📝 Description

This is a simple API to manage your TODO list.

Written in Go, with a PostgreSQL database.

## 🚀 Installation and usage with Docker Compose

### 📦 Prerequisites

- Docker
- Docker Compose

### 📦 Installation

```bash
git clone 
cd todoom
docker-compose up -d
```
### 📦 Run the SQL script

```bash
docker exec -i db_api_todo psql -U postgres -d db_api_todo < ./db/database.sql
```

### 📦 Usage

Routes:

- [GET] /todos

Get one todo by id

- [GET] /todos/list

Get all todos

- [POST] /todos/create

Create a new todo

- [PUT] /todos/update

Update a todo

- [DELETE] /todos/remove

Delete a todo