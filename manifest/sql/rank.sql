CREATE TABLE semester_rank
(
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,
    stu_id                  INTEGER   NOT NULL, -- 學生_id
    year                    INTEGER   NOT NULL, -- 學年
    semester                INTEGER   NOT NULL, -- 學期
    total_points            DOUBLE    NOT NULL, -- 總積分
    total_credits           INTEGER   NOT NULL, -- 總修學分數
    earned_credits          INTEGER   NOT NULL, -- 實得學分數
    failed_credits          INTEGER   NOT NULL, -- 不及格學分數
    average_score           DOUBLE    NOT NULL, -- 平均成績
    department_rank         INTEGER   NOT NULL, -- 系排名
    department_rank_percent DOUBLE    NOT NULL, -- 系排名百分比
    current_student_status  TEXT      NOT NULL, -- 目前學籍狀態
    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP,

    FOREIGN KEY (stu_id) REFERENCES students (id),
    UNIQUE (stu_id, year, semester)
);