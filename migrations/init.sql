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
    "responsible" uuid NOT NULL REFERENCES "Department" ("id")
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
    "responsible_id" uuid NOT NULL REFERENCES "Employee" ("id"),
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
