FROM redis:latest
COPY redis.conf /usr/local/etc/redis/redis.conf
# Chạy redis-server kèm đường dẫn file cấu hình
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]