provider "redis" {
  address  = "localhost:6379"
  password = "your_redis_password"
}

resource "redis_user" "example" {
  username = "myuser"
  password = "mypass123"
  commands = ["get", "set", "del"]
  keys     = ["*"]
  enabled  = true
}