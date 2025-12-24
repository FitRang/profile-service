# Profile Service

The **Profile Service** is responsible for managing all **Profile** and **Dossier** related logic in the FitRang ecosystem. It acts as the single source of truth for user identity-related metadata beyond authentication.

This service is written in **Go** and primarily communicates using **GraphQL**, while also exposing a **gRPC interface** for internal service-to-service communication.

---

## Responsibilities

### Profile Management

* Create user profiles using **username and email**
* Edit profile details
* Delete profiles
* Fetch profile data

### Dossier Management

* Create dossiers linked to a **username and email**
* Update dossier information
* Delete dossiers
* Fetch dossier details

> A **Profile** represents basic user information, while a **Dossier** represents deeper personal attributes used for styling, recommendations, and personalization.

---

## Communication Protocols

### GraphQL (Primary)

* Main API exposed to frontend and external consumers
* Used for:

  * Profile creation & updates
  * Dossier management
  * Querying profile and dossier data

### gRPC (Internal)

* Used by other backend services
* Provides high-performance access to:

  * Profile data
  * Dossier data
* Identifies users using **username and email**

---

## User Identification Model

* The Profile Service **does not use `owner_id`**
* Users are uniquely identified by:

  * `username`
  * `email`
* These fields are treated as trusted identifiers coming from the authentication layer (Firebase)

---

## Tech Stack

* **Language:** Go
* **API Protocols:**

  * GraphQL (primary)
  * gRPC (internal)
* **Architecture:** Microservice
* **Deployment:** Containerized (Docker), Kubernetes-ready

---

## High-Level Flow

1. User authenticates via **Firebase**
2. Firebase issues identity claims (username & email)
3. Frontend interacts with Profile Service via **GraphQL**
4. Profile Service:

   * Creates or updates profile and dossier data using username & email
5. Other backend services fetch profile/dossier data via **gRPC**

---

## Features

* Full CRUD support for Profiles
* Full CRUD support for Dossiers
* GraphQL schema-first API design
* gRPC endpoints for internal data access
* Username & email–based identification
* Designed for scalability and extensibility

---

## TODOs

* [x] Create Profile API
* [x] Edit Profile API
* [x] Delete Profile API
* [x] Create Dossier API
* [x] Edit Dossier API
* [x] Delete Dossier API
* [x] GraphQL schema and resolvers
* [x] gRPC server setup
* [x] gRPC endpoints for profile retrieval
* [x] gRPC endpoints for dossier retrieval
* [x] Authorization integration
* [ ] Test upload endpoint
* [ ] Add check username endpoint

---

## Future Improvements

* Username availability checks
* Media upload validations
* Caching frequently accessed profiles
* Rate limiting on profile updates
* Audit logs for profile changes

---

## Notes

* This service **does not handle authentication**
* Authentication and identity verification are handled by **Firebase**
* GraphQL is the public contract; gRPC is an internal optimization layer
* Username and email are considered immutable identifiers once created
