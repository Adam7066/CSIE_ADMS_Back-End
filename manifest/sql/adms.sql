CREATE TABLE star_plan -- 繁星計畫
(
    id                               INTEGER PRIMARY KEY AUTOINCREMENT,
    stu_id                           INTEGER NOT NULL, -- 學生_id
    gsat_score_id                    INTEGER,          -- 學測成績_id
    school_dept_code                 TEXT,             -- 校系代碼
    recommended_order                INTEGER,          -- 推薦順位
    adms_order                       INTEGER,          -- 錄取(通過篩選)志願序
    school_ranking_percentage        INTEGER,          -- 在校學業成績全校排名百分比
    chinese_ranking_percentage       INTEGER,          -- 國文學業總平均成績全校排名百分比
    english_ranking_percentage       INTEGER,          -- 英文學業總平均成績全校排名百分比
    math_ranking_percentage          INTEGER,          -- 數學學業總平均成績全校排名百分比
    physics_ranking_percentage       INTEGER,          -- 物理學業總平均成績全校排名百分比
    chemistry_ranking_percentage     INTEGER,          -- 化學學業總平均成績全校排名百分比
    biology_ranking_percentage       INTEGER,          -- 生物學業總平均成績全校排名百分比
    earth_science_ranking_percentage INTEGER,          -- 地球科學學業總平均成績全校排名百分比
    history_ranking_percentage       INTEGER,          -- 歷史學業總平均成績全校排名百分比
    geography_ranking_percentage     INTEGER,          -- 地理學業總平均成績全校排名百分比
    citizenship_ranking_percentage   INTEGER,          -- 公民與社會學業總平均成績全校排名百分比
    music_ranking_percentage         INTEGER,          -- 音樂學業總平均成績全校排名百分比
    art_ranking_percentage           INTEGER,          -- 美術學業總平均成績全校排名百分比
    dance_ranking_percentage         INTEGER,          -- 舞蹈學業總平均成績全校排名百分比
    physical_ranking_percentage      INTEGER,          -- 體育學業總平均成績全校排名百分比
    study_subject_course_class       TEXT,             -- 就讀科、學程、班別
    FOREIGN KEY (stu_id) REFERENCES students (id),
    FOREIGN KEY (gsat_score_id) REFERENCES gsat_score (id)
);

CREATE TABLE apply_general -- 申請入學_一般組
(
    id                    INTEGER PRIMARY KEY AUTOINCREMENT,
    stu_id                INTEGER NOT NULL, -- 學生_id
    gsat_score_id         INTEGER,          -- 學測成績_id
    school_dept_code      TEXT,             -- 校系代碼
    gsat_cal_score        DOUBLE,           -- 學科能力測驗成績
    project_score_1       DOUBLE,           -- 指定項目一成績
    project_test_score    DOUBLE,           -- 指定項目甄試成績
    selection_total_score DOUBLE,           -- 甄選總成績
    adms_status           TEXT,             -- 招生名額錄取狀態
    adms_rank             INTEGER,          -- 招生名額名次
    adms_total_rank       INTEGER,          -- 招生名額總名次
    FOREIGN KEY (stu_id) REFERENCES students (id),
    FOREIGN KEY (gsat_score_id) REFERENCES gsat_score (id)
);

CREATE TABLE apply_apcs -- 申請入學_APCS
(
    id                    INTEGER PRIMARY KEY AUTOINCREMENT,
    stu_id                INTEGER NOT NULL, -- 學生_id
    gsat_score_id         INTEGER,          -- 學測成績_id
    school_dept_code      TEXT,             -- 校系代碼
    gsat_cal_score        DOUBLE,           -- 學科能力測驗成績
    project_score_1       DOUBLE,           -- 指定項目一成績
    project_score_2       DOUBLE,           -- 指定項目二成績
    project_test_score    DOUBLE,           -- 指定項目甄試成績
    selection_total_score DOUBLE,           -- 甄選總成績
    adms_status           TEXT,             -- 招生名額錄取狀態
    adms_rank             INTEGER,          -- 招生名額名次
    adms_total_rank       INTEGER,          -- 招生名額總名次
    FOREIGN KEY (stu_id) REFERENCES students (id),
    FOREIGN KEY (gsat_score_id) REFERENCES gsat_score (id)
);

CREATE TABLE gsat_score -- 學測成績
(
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    gsat_test_number TEXT,                                                    -- 學測應試號碼
    chinese          INTEGER CHECK ( chinese BETWEEN 0 AND 15),               -- 國文級分
    english          INTEGER CHECK ( english BETWEEN 0 AND 15),               -- 英文級分
    math             INTEGER CHECK ( math BETWEEN 0 AND 15),                  -- 數學級分
    nature           INTEGER CHECK ( nature BETWEEN 0 AND 15),                -- 自然級分
    society          INTEGER CHECK ( society BETWEEN 0 AND 15),               -- 社會級分
    listening        TEXT CHECK ( listening IN ('A', 'B', 'C', 'F', '缺考')), -- 聽力測驗
    sum_score        INTEGER CHECK ( sum_score BETWEEN 0 AND 75)              -- 總級分
);

CREATE TRIGGER update_sum_score
    AFTER INSERT
    ON gsat_score
BEGIN
    UPDATE gsat_score
    SET sum_score = chinese + english + math + nature + society
    WHERE id = NEW.id;
END;

CREATE TRIGGER update_sum_after_update
    AFTER UPDATE OF chinese, english, math, nature, society
    ON gsat_score
BEGIN
    UPDATE gsat_score
    SET sum_score = chinese + english + math + nature + society
    WHERE id = NEW.id;
END;
