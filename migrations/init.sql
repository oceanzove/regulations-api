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
                              "account_email" VARCHAR NOT NULL REFERENCES "Account"("email"),
                              "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                              "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Таблица с процессами
CREATE TABLE "Process" (
                           "id" SERIAL PRIMARY KEY,
                           "title" VARCHAR NOT NULL,
                           "description" TEXT,
                           "regulation_id" INT REFERENCES "Regulation"("id"),
                           "account_email" VARCHAR NOT NULL REFERENCES "Account"("email"),
                           "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Таблица с шагами для процессов
CREATE TABLE "Step" (
                        "id" SERIAL PRIMARY KEY,
                        "name" VARCHAR NOT NULL,
                        "description" TEXT,
                        "process_id" INT REFERENCES "Process"("id") ON DELETE CASCADE,
                        "order" INT NOT NULL CHECK ("order" > 0)
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

CREATE TRIGGER update_process_updated_at
    BEFORE UPDATE ON "Process"
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

-- Вставляем данные в таблицу Process
INSERT INTO "Process" ("title", "description", "regulation_id", "account_email") VALUES
                                                                   ('Процесс 1', 'Описание процесса 1', 1, 'test'),
                                                                   ('Процесс 2', 'Описание процесса 2', 2, 'test');

-- Вставляем данные в таблицу Step
INSERT INTO "Step" ("name", "description", "process_id", "order") VALUES
                                                                      ('Шаг 1', 'Описание шага 1', 1, 1),
                                                                      ('Шаг 2', 'Описание шага 2', 1, 2),
                                                                      ('Шаг 1', 'Описание шага 1 для процесса 2', 2, 1);

-- Индексы для повышения производительности
CREATE INDEX idx_account_email ON "Regulation" ("account_email");
CREATE INDEX idx_process_regulation_id ON "Process" ("regulation_id");
CREATE INDEX idx_step_process_id ON "Step" ("process_id");
