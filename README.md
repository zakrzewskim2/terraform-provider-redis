# Terraform Provider: Redis

A Terraform provider to manage [Redis](https://redis.io) users using Redis' ACL commands (`ACL SETUSER`, `ACL DELUSER`, etc.).  
Supports creating, updating, and deleting Redis users.

## ⚡ Features

- Create Redis users
- Set passwords, allowed commands, and key access
- Enable or disable users
- Compatible with Redis 6 and later

---

## 🧰 Requirements

- Terraform >= 1.0.0
- Redis >= 6.0 (for ACL support)
- Go >= 1.19 (for building the provider)

---

## 📦 Installation

To use this provider, add it to your Terraform configuration:

```hcl
terraform {
  required_providers {
    redis = {
      source  = "zakrzewskim2/redis"
      version = "~> 0.1.0"
    }
  }
}

provider "redis" {
  address  = "localhost:6379"
  password = "your_redis_password" # Optional
}
