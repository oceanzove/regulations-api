### Конфигурационный файл (`configs/config.json`)

Конфигурационный файл (`config.json`) используется для предоставления настроек сервера, когда `IsReadConfig`
установлен в `true`. Ниже приведен пример структуры конфигурации:

````json
{
  "server": {
    "server_port": "25504",
    "jwt_secret_key": "key"
  },
  "regulations-database": {
    "db_host": "localhost",
    "db_port": "5433",
    "db_username": "postgres",
    "db_name": "postgres",
    "db_ssl_mode": "disable"
  }
}
```