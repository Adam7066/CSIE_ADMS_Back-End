CREATE TABLE students
(
    id                   INTEGER PRIMARY KEY AUTOINCREMENT,
    name                 TEXT      NOT NULL, -- 姓名
    gender_id            INTEGER,            -- 性別_id
    student_code         TEXT      NOT NULL, -- 學號
    degree_id            INTEGER   NOT NULL, -- 學歷_id
    admission_method_id  INTEGER   NOT NULL, -- 入學方式_id
    admission_group_id   INTEGER,            -- 入學組別_id
    identity_category_id INTEGER,            -- 身分類別_id
    graduated_school_id  INTEGER   NOT NULL, -- 畢業學校_id
    created_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at           TIMESTAMP,

    FOREIGN KEY (gender_id) REFERENCES genders (id),
    FOREIGN KEY (degree_id) REFERENCES degrees (id),
    FOREIGN KEY (admission_method_id) REFERENCES admission_methods (id),
    FOREIGN KEY (admission_group_id) REFERENCES admission_groups (id),
    FOREIGN KEY (identity_category_id) REFERENCES identity_categories (id),
    FOREIGN KEY (graduated_school_id) REFERENCES schools (id)
);

CREATE TABLE genders
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 性別名稱
);

CREATE TABLE degrees
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 學歷名稱
);

CREATE TABLE admission_methods
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 入學方式名稱
);

CREATE TABLE admission_groups
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 入學組別名稱
);

CREATE TABLE identity_categories
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 身分類別名稱
);

CREATE TABLE schools
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    degree_id        NOT NULL, -- 學歷_id
    name        TEXT NOT NULL, -- 名稱
    school_code TEXT NOT NULL, -- 代碼
    city_id          NOT NULL, -- 所在縣市_id
    area_id          NOT NULL, -- 所在區域_id

    FOREIGN KEY (degree_id) REFERENCES degrees (id),
    FOREIGN KEY (city_id) REFERENCES cities (id),
    FOREIGN KEY (area_id) REFERENCES areas (id)
);

CREATE TABLE cities
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 縣市名稱
);

CREATE TABLE areas
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL -- 區域名稱
);