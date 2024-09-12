# Stage 1: Build the app
FROM node:16-alpine AS build-stage

WORKDIR /app

# Устанавливаем зависимости
COPY package*.json ./
RUN npm install

# Копируем исходный код
COPY . .

# Сборка приложения
RUN npm run build

# Stage 2: Используем Nginx для сервинга
FROM nginx:alpine

# Копируем сгенерированные файлы Vue.js
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Копируем скрипт для динамической замены переменных
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Копируем шаблон конфигурации
COPY ./public/config-template.js /usr/share/nginx/html/

# Открываем порт
EXPOSE 80

# Запуск скрипта entrypoint
ENTRYPOINT ["/entrypoint.sh"]