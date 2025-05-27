CREATE TYPE role_enum AS ENUM ('Admin', 'User', 'Guest');

-- Таблица пользователей
CREATE TABLE "Account"
(
    "id"       uuid PRIMARY KEY,
    "email"    VARCHAR   NOT NULL,
    "password" VARCHAR   NOT NULL,
    "role"     role_enum NOT NULL
);

-- Каталог всех секций (Section)
CREATE TABLE "Section"
(
    "id"         uuid PRIMARY KEY,
    "title"      VARCHAR NOT NULL,
    "content"    TEXT,
    "account_id" uuid    NOT NULL REFERENCES "Account" ("id"),
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Таблица с регламентами
CREATE TABLE "Regulation"
(
    "id"         uuid PRIMARY KEY,
    "title"      VARCHAR NOT NULL,
    "content"    TEXT,
    "account_id" uuid    NOT NULL REFERENCES "Account" ("id"),
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Связь: какие секции в каком регламенте (RegulationSection)
CREATE TABLE "RegulationSection"
(
    "id"            uuid PRIMARY KEY,
    "regulation_id" uuid REFERENCES "Regulation" ("id") ON DELETE CASCADE,
    "section_id"    uuid REFERENCES "Section" ("id"),
    "order"         INT NOT NULL
);

-- Таблица с процессами (убрали regulation_id!)
CREATE TABLE "Process"
(
    "id"          uuid PRIMARY KEY,
    "title"       VARCHAR NOT NULL,
    "description" TEXT,
    "responsible" VARCHAR DEFAULT '',
    "account_id"  uuid    NOT NULL REFERENCES "Account" ("id"),
    "created_at"  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Таблица-связка многие-ко-многим
CREATE TABLE "ProcessRegulation"
(
    "process_id"    uuid REFERENCES "Process" ("id") ON DELETE CASCADE,
    "regulation_id" uuid REFERENCES "Regulation" ("id") ON DELETE CASCADE,
    PRIMARY KEY ("process_id", "regulation_id")
);

-- Таблица с шагами для процессов
CREATE TABLE "Step"
(
    "id"          uuid PRIMARY KEY,
    "name"        VARCHAR NOT NULL,
    "description" TEXT,
    "responsible" VARCHAR,
    "process_id"  uuid REFERENCES "Process" ("id") ON DELETE CASCADE,
    "order"       INT     NOT NULL
);

-- Аудит (кто, когда, что редактировал)
CREATE TABLE "RegulationAudit"
(
    "id"            uuid PRIMARY KEY,
    "regulation_id" uuid REFERENCES "Regulation" ("id"),
    "changed_by"    uuid REFERENCES "Account" ("id"),
    "change_type"   VARCHAR NOT NULL,
    "change_time"   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "old_value"     TEXT,
    "new_value"     TEXT
);

-- Триггер для автоматического обновления столбца updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_regulation_updated_at
    BEFORE UPDATE
    ON "Regulation"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_process_updated_at
    BEFORE UPDATE
    ON "Process"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Вставляем данные в таблицу Account
INSERT INTO "Account" ("id", "email", "password", "role")
VALUES ('2497a896-7e45-4f53-b7c0-318df7569c75', 'test', '$2a$10$A6lwFl3TlDyZA3UkbobGFOFe.PyI0iQsEUlm1rILs7HWGAaV3AKcC',
        'Admin');

-- Вставляем данные в таблицу Regulation
INSERT INTO "Regulation" ("id", "title", "content", "account_id")
VALUES (gen_random_uuid(), 'Регламент 1', 'Содержимое регламента 1',
        '2497a896-7e45-4f53-b7c0-318df7569c75'),
       (gen_random_uuid(), 'Регламент 2', 'Содержимое регламента 2',
        '2497a896-7e45-4f53-b7c0-318df7569c75');

-- Вставляем данные в таблицу Process
INSERT INTO "Process" ("id", "title", "description", "account_id")
VALUES (gen_random_uuid(), 'Процесс 1', 'Описание процесса 1', '2497a896-7e45-4f53-b7c0-318df7569c75'),
       (gen_random_uuid(), 'Процесс 2', 'Описание процесса 2', '2497a896-7e45-4f53-b7c0-318df7569c75');

-- -- Привязываем процессы к регламентам (многие-ко-многим)
-- -- Предположим, что у процессов id = 'id1', 'id2' (замени на реальные uuid)
-- INSERT INTO "ProcessRegulation" ("process_id", "regulation_id")
-- VALUES ('id1', 1),
--        ('id1', 2), -- Процесс 1 относится к двум регламентам
--        ('id2', 1);
-- Процесс 2 к первому регламенту

-- Вставляем данные в таблицу Step
-- INSERT INTO "Step" ("id", "name", "description", "process_id", "order")
-- VALUES (gen_random_uuid(), 'Шаг 1', 'Описание шага 1', 'id1', 1),
--        (gen_random_uuid(), 'Шаг 2', 'Описание шага 2', 'id1', 2),
--        (gen_random_uuid(), 'Шаг 1', 'Описание шага 1 для процесса 2', 'id2', 1);

-- Индексы для повышения производительности
CREATE INDEX idx_account_id ON "Regulation" ("id");
CREATE INDEX idx_step_process_id ON "Step" ("process_id");
CREATE INDEX idx_processregulation_process_id ON "ProcessRegulation" ("process_id");
CREATE INDEX idx_processregulation_regulation_id ON "ProcessRegulation" ("regulation_id");