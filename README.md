# REST API

A web server that serves data, recieves data and stores data and can be used as a backend for a mobile app or website.

- **GET** /events

  - `get a list of available events`

- **GET** /events/{id}
  - `get an event by id`
- **POST** /events
  - `create a new bookable event`
- **PUT** /events/{id}
  - `update an event`
- **DELETE** /events/{id}
  - `delete an event`
- **POST** /signup
  - `create a new user`
- **POST** /login
  - `authenticate user and create session`
- **POST** /events/{id}/register
  - `register user for an event`
- **DELETE** /events/{id}/register
  - `cancel event registration`
