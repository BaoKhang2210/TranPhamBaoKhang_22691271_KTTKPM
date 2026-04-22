FROM postgres:15
# Postgres sẽ tự động chạy mọi file .sql trong thư mục này khi khởi tạo lần đầu
COPY init.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_PASSWORD=mysecretpassword