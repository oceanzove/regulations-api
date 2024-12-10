-- Таблица пользователей
CREATE TABLE "Role" (
                        "name" VARCHAR PRIMARY KEY
);

CREATE TABLE "Account" (
                           "email" VARCHAR PRIMARY KEY,
                           "password" VARCHAR NOT NULL,
                           "role" VARCHAR NOT NULL REFERENCES "Role"("name")
);

-- Вставляем данные в таблицу Role
INSERT INTO "Role" ("name") VALUES
                                ('Admin'),
                                ('User'),
                                ('Guest');

-- Вставляем данные в таблицу Account
INSERT INTO "Account" ("email", "password", "role") VALUES
                                ('zove', '$2a$10$QkjvoLbAM3bDlCgDqu/G4eMfdu0FcLAPSXj4OjwKBRXC79jiJaMtO', 'Admin');


-- Индексы для повышения производительности
-- CREATE INDEX idx_process_id ON regulations (process_id);
-- CREATE INDEX idx_regulation_id ON sections (regulation_id);
-- CREATE INDEX idx_regulation_view_regulation_id ON regulation_views (regulation_id);
-- CREATE INDEX idx_username ON users (username);
-- CREATE INDEX idx_email ON users (email);
