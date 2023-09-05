FROM postgres:10.0-alpine

ENV POSTGRES_USER admin
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB testdb

# Открытие порта для доступа к PostgreSQL
EXPOSE 5432

# Запуск PostgreSQL при запуске контейнера
CMD ["postgres"]
