import os
env = os.getenv("APP_ENV", "production")
print(f"Ứng dụng đang chạy trong môi trường: {env}")