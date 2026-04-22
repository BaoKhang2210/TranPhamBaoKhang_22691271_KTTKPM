# Sử dụng image Nginx mới nhất
FROM nginx:latest

# Copy file index.html từ máy cậu vào đúng thư mục mặc định của Nginx trong container
COPY index.html /usr/share/nginx/html/index.html

# Nginx chạy mặc định ở cổng 80, ta EXPOSE để đánh dấu
EXPOSE 80