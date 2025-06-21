CREATE TYPE role_enum AS ENUM ('administrator', 'employee');
CREATE TYPE marital_enum AS ENUM ('single', 'married', 'widowed', 'divorced');

-- Таблица организаций
CREATE TABLE IF NOT EXISTS "Organization"
(
    "id"   uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL UNIQUE
);

-- Таблица отделов
CREATE TABLE IF NOT EXISTS "Department"
(
    "id"              uuid PRIMARY KEY,
    "name"            varchar(255) NOT NULL UNIQUE,
    "organization_id" uuid         NOT NULL REFERENCES "Organization" ("id")
);



-- Таблица сотрудников
CREATE TABLE IF NOT EXISTS "Employee"
(
    "id"                  uuid PRIMARY KEY,
    "full_name"           varchar(255) NOT NULL,
    "phone_number"        varchar(255) NOT NULL,
    "birth_date"          date         NOT NULL,
    "employment_date"     date         NOT NULL,
    "residential_address" varchar(255) NOT NULL,
    "marital_status"      marital_enum NOT NULL,
    "email"               varchar(255) NOT NULL
);

-- Таблица должностей
CREATE TABLE IF NOT EXISTS "Position"
(
    "id"   uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL
);

-- Таблица связи сотрудника с должностью
CREATE TABLE IF NOT EXISTS "EmployeePosition"
(
    "employee_id" uuid REFERENCES "Employee" ("id") ON DELETE CASCADE,
    "position_id" uuid REFERENCES "Position" ("id"),
    PRIMARY KEY ("employee_id", "position_id")
);

-- Таблица связи отдела с должностями
CREATE TABLE IF NOT EXISTS "DepartmentPosition"
(
    "department_id" uuid REFERENCES "Department" ("id"),
    "position_id"   uuid REFERENCES "Position" ("id"),
    PRIMARY KEY ("department_id", "position_id")
);

-- Таблица связи сотрудника с отделом
CREATE TABLE IF NOT EXISTS "EmployeeDepartment"
(
    "employee_id"   uuid REFERENCES "Employee" ("id"),
    "department_id" uuid REFERENCES "Department" ("id"),
    PRIMARY KEY ("employee_id", "department_id")
);

-- Таблица пользователей
CREATE TABLE "Account"
(
    "id"       uuid PRIMARY KEY REFERENCES "Employee" ("id"),
    "login"    VARCHAR   NOT NULL,
    "password" VARCHAR   NOT NULL,
    "role"     role_enum NOT NULL
);

-- Каталог всех секций (Section)
CREATE TABLE "Section"
(
    "id"            uuid PRIMARY KEY,
    "title"         VARCHAR NOT NULL,
    "content"       TEXT,
    "department_id" uuid    NOT NULL REFERENCES "Department" ("id")
);

-- Таблица с регламентами
CREATE TABLE "Regulation"
(
    "id"            uuid PRIMARY KEY,
    "title"         VARCHAR NOT NULL,
    "content"       TEXT,
    "department_id" uuid    NOT NULL REFERENCES "Department" ("id")
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
    "responsible" uuid    NOT NULL REFERENCES "Department" ("id")
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
    "id"             uuid PRIMARY KEY,
    "name"           VARCHAR NOT NULL,
    "description"    TEXT,
    "responsible_id" uuid    NOT NULL REFERENCES "Position" ("id"),
    "process_id"     uuid REFERENCES "Process" ("id") ON DELETE CASCADE,
    "order"          INT     NOT NULL
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

-- Аудит (кто, когда, что редактировал)
CREATE TABLE "ProcessAudit"
(
    "id"            uuid PRIMARY KEY,
    "regulation_id" uuid REFERENCES "Process" ("id"),
    "changed_by"    uuid REFERENCES "Account" ("id"),
    "change_type"   VARCHAR NOT NULL,
    "change_time"   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "old_value"     TEXT,
    "new_value"     TEXT
);

-- Организации
INSERT INTO "Organization" ("id", "name")
VALUES ('a1111111-1111-1111-1111-111111111111', 'ООО СтартСет');

-- Отделы
INSERT INTO "Department" ("id", "name", "organization_id")
VALUES ('d1111111-1111-1111-1111-111111111111', 'Отдел разработки', 'a1111111-1111-1111-1111-111111111111'),
       ('d2222222-2222-2222-2222-222222222222', 'Отдел кадров', 'a1111111-1111-1111-1111-111111111111'),
       ('d3333333-3333-3333-3333-333333333333', 'Маркетинг', 'a1111111-1111-1111-1111-111111111111');

-- Должности
INSERT INTO "Position" ("id", "name")
VALUES ('b1111111-1111-1111-1111-111111111111', 'Разработчик'),
       ('b2222222-2222-2222-2222-222222222222', 'Тестировщик'),
       ('b3333333-3333-3333-3333-333333333333', 'HR'),
       ('b4444444-4444-4444-4444-444444444444', 'Маркетолог');

-- Связь отделов с должностями
INSERT INTO "DepartmentPosition" ("department_id", "position_id")
VALUES ('d1111111-1111-1111-1111-111111111111', 'b1111111-1111-1111-1111-111111111111'), -- Разработка - разработчик
       ('d1111111-1111-1111-1111-111111111111', 'b2222222-2222-2222-2222-222222222222'), -- Разработка - тестировщик
       ('d2222222-2222-2222-2222-222222222222', 'b3333333-3333-3333-3333-333333333333'), -- Кадры - HR
       ('d3333333-3333-3333-3333-333333333333', 'b4444444-4444-4444-4444-444444444444');

-- Сотрудники
INSERT INTO "Employee" ("id", "full_name", "phone_number", "birth_date", "employment_date", "residential_address",
                        "marital_status", "email")
VALUES ('e1111111-1111-1111-1111-111111111111', 'Иванов Иван Иванович', '+70000000001', '1990-01-01', '2020-01-10',
        'г. Москва, ул. Ленина, д.1', 'single', 'ivanov@example.com'),
       ('e2222222-2222-2222-2222-222222222222', 'Петров Петр Петрович', '+70000000002', '1985-05-12', '2019-03-15',
        'г. Москва, ул. Гагарина, д.5', 'married', 'petrov@example.com'),
       ('e3333333-3333-3333-3333-333333333333', 'Сидорова Анна Сергеевна', '+70000000003', '1992-09-20', '2022-06-01',
        'г. СПб, пр. Невский, д.10', 'single', 'sidorova@example.com');

-- Связь сотрудник-должность
INSERT INTO "EmployeePosition" ("employee_id", "position_id")
VALUES ('e1111111-1111-1111-1111-111111111111', 'b1111111-1111-1111-1111-111111111111'), -- Иванов - разработчик
       ('e2222222-2222-2222-2222-222222222222', 'b2222222-2222-2222-2222-222222222222'), -- Петров - тестировщик
       ('e3333333-3333-3333-3333-333333333333', 'b3333333-3333-3333-3333-333333333333');

-- Связь сотрудник-отдел
INSERT INTO "EmployeeDepartment" ("employee_id", "department_id")
VALUES ('e1111111-1111-1111-1111-111111111111', 'd1111111-1111-1111-1111-111111111111'), -- Иванов - разработка
       ('e2222222-2222-2222-2222-222222222222', 'd1111111-1111-1111-1111-111111111111'), -- Петров - разработка
       ('e3333333-3333-3333-3333-333333333333', 'd2222222-2222-2222-2222-222222222222');

-- Пользовательские аккаунты (id = id сотрудника!)
INSERT INTO "Account" ("id", "login", "password", "role")
VALUES ('e1111111-1111-1111-1111-111111111111', 'ivanov',
        '$2a$10$A6lwFl3TlDyZA3UkbobGFOFe.PyI0iQsEUlm1rILs7HWGAaV3AKcC', 'employee'),
       ('e2222222-2222-2222-2222-222222222222', 'petrov',
        '$2a$10$A6lwFl3TlDyZA3UkbobGFOFe.PyI0iQsEUlm1rILs7HWGAaV3AKcC', 'administrator'),
       ('e3333333-3333-3333-3333-333333333333', 'sidorova',
        '$2a$10$A6lwFl3TlDyZA3UkbobGFOFe.PyI0iQsEUlm1rILs7HWGAaV3AKcC', 'employee');


--- Каталог секций
INSERT INTO "Section" ("id", "title", "content", "department_id")
VALUES ('c1111111-1111-1111-1111-111111111111', 'Безопасность', 'Описание политики безопасности',
        'd1111111-1111-1111-1111-111111111111'),
       ('c2222222-2222-2222-2222-222222222222', 'Отчётность', 'Правила по отчётности',
        'd2222222-2222-2222-2222-222222222222');

-- Регламенты
-- INSERT INTO "Regulation" ("id", "title", "content", "department_id")
-- VALUES ('eg111111-1111-1111-1111-111111111111', 'Регламент безопасности', 'Текст регламента',
--         'd1111111-1111-1111-1111-111111111111'),
--        ('eg222222-2222-2222-2222-222222222222', 'Регламент отчётности', 'Текст отчётности',
--         'd2222222-2222-2222-2222-222222222222');
--
-- -- Связь секций и регламентов (порядок)
-- INSERT INTO "RegulationSection" ("id", "regulation_id", "section_id", "order")
-- VALUES ('rs1111111-1111-1111-1111-111111111111', 'r1111111-1111-1111-1111-111111111111',
--         'c1111111-1111-1111-1111-111111111111', 1),
--        ('rs2222222-2222-2222-2222-222222222222', 'r2222222-2222-2222-2222-222222222222',
--         'c2222222-2222-2222-2222-222222222222', 1);
--
-- -- Процессы
-- INSERT INTO "Process" ("id", "title", "description", "responsible")
-- VALUES ('pr111111-1111-1111-1111-111111111111', 'Процесс внедрения', 'Описание внедрения',
--         'd1111111-1111-1111-1111-111111111111'),
--        ('pr222222-2222-2222-2222-222222222222', 'Процесс отчетности', 'Описание отчетности',
--         'd2222222-2222-2222-2222-222222222222');
--
-- -- Связь процессы-регламенты
-- INSERT INTO "ProcessRegulation" ("process_id", "regulation_id")
-- VALUES ('pr111111-1111-1111-1111-111111111111', 'r1111111-1111-1111-1111-111111111111'),
--        ('pr222222-2222-2222-2222-222222222222', 'r2222222-2222-2222-2222-222222222222');
--
-- -- Шаги процесса
-- INSERT INTO "Step" ("id", "name", "description", "responsible_id", "process_id", "order")
-- VALUES ('st111111-1111-1111-1111-111111111111', 'Шаг 1', 'Первый шаг процесса', 'e1111111-1111-1111-1111-111111111111',
--         'pr111111-1111-1111-1111-111111111111', 1),
--        ('st222222-2222-2222-2222-222222222222', 'Шаг 2', 'Второй шаг процесса', 'e2222222-2222-2222-2222-222222222222',
--         'pr222222-2222-2222-2222-222222222222', 1);
--
-- -- Аудит регламентов
-- INSERT INTO "RegulationAudit" ("id", "regulation_id", "changed_by", "change_type", "old_value", "new_value")
-- VALUES ('a1111111-1111-1111-1111-111111111111', 'r1111111-1111-1111-1111-111111111111',
--         'e2222222-2222-2222-2222-222222222222', 'update', 'старый текст', 'новый текст');
--
-- -- Аудит процессов
-- INSERT INTO "ProcessAudit" ("id", "regulation_id", "changed_by", "change_type", "old_value", "new_value")
-- VALUES ('a2222222-2222-2222-2222-222222222222', 'pr111111-1111-1111-1111-111111111111',
--         'e1111111-1111-1111-1111-111111111111', 'create', NULL, 'создан процесс');