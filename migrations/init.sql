-- Таблица пользователей
CREATE TABLE "Role" (
                        "name" VARCHAR PRIMARY KEY
);

CREATE TABLE "Account" (
                           "email" VARCHAR PRIMARY KEY,
                           "password" VARCHAR NOT NULL,
                           "role" VARCHAR NOT NULL REFERENCES "Role"("name")
);

-- Таблица с регламентами
CREATE TABLE "Regulation" (
                              "id" SERIAL PRIMARY KEY,
                              "title" VARCHAR NOT NULL,
                              "content" TEXT,
                              "account_email" VARCHAR NOT NULL REFERENCES "Account"("email") ON DELETE CASCADE,
                              "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                              "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Триггер для автоматического обновления столбца updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_regulation_updated_at
    BEFORE UPDATE ON "Regulation"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Вставляем данные в таблицу Role
INSERT INTO "Role" ("name") VALUES
                                ('Admin'),
                                ('User'),
                                ('Guest');

-- Вставляем данные в таблицу Account
INSERT INTO "Account" ("email", "password", "role") VALUES
                                ('test', '$2a$10$A6lwFl3TlDyZA3UkbobGFOFe.PyI0iQsEUlm1rILs7HWGAaV3AKcC', 'Admin');

-- Вставляем данные в таблицу Regulation
INSERT INTO "Regulation" ("title", "content", "account_email") VALUES
                                                                   ('Регламент 1', 'Содержимое регламента 1', 'test'),
                                                                   ('Регламент 2', 'Содержимое регламента 2', 'test');

-- Индексы для повышения производительности
CREATE INDEX idx_account_email ON "Regulation" ("account_email");

-- CREATE INDEX idx_process_id ON regulations (process_id);
-- CREATE INDEX idx_regulation_id ON sections (regulation_id);
-- CREATE INDEX idx_regulation_view_regulation_id ON regulation_views (regulation_id);
-- CREATE INDEX idx_username ON users (username);
-- CREATE INDEX idx_email ON users (email);
