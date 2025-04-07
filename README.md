# 👤 Profile Service - FitRang

The **Profile Service** is a core microservice in the FitRang platform. It stores and manages additional user metadata that is not handled by the authentication provider (Keycloak). This includes subscription plan info, trial status, user preferences, and profile visibility settings.

---

## 🧠 Responsibilities

- Store user-specific data such as:
  - Subscription type (Free, Trial, Premium)
  - Trial status and duration
  - User role (User, Stylist, Admin)
  - Profile visibility (Public/Private)
- Sync user data on new registration (triggered by Keycloak)
- Provide user data to other services (via gRPC)

---

## ⚙️ Tech Stack

- **Language**: GO 
- **Framework**: Gin 
- **Database**: Supabase DB
- **Deployment**: Docker-ready, deployable on Kubernetes or Lambda

---

## 📁 Folder Structure
## 📦 API Endpoints
## 🛠️ Local Development
## 📄 License

MIT License

---

## 👥 Organization

Maintained by [Foxtrot](https://github.com/Foxtrot-14) and other contributors.

