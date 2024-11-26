create table if not exists startUp (
                         id SERIAL PRIMARY KEY,
                         userId int not null,

    -- Информация о стартапе
                         title VARCHAR(100) NOT NULL,
                         topic VARCHAR(100) NOT NULL,
                         idea TEXT DEFAULT '',
                         strategy TEXT DEFAULT '',
                         historyOfCreation TEXT DEFAULT '',
                         status smallint NOT NULL DEFAULT 0,
                         stage smallint NOT NULL DEFAULT 0,

    -- Цель по финансированию, сколько нужно собрать
                         fundingGoal NUMERIC(15, 5) DEFAULT 0.0,

    -- Предлагаемый процент инвестору
                         offeredPercent NUMERIC(8, 5) DEFAULT 0.0,

    -- Информация о создателе
                         founderFullName VARCHAR(255) NOT NULL,
                         founderEmail VARCHAR(255) NOT NULL UNIQUE,
                         founderSocials jsonb,

    -- Информация о дате создания и обновления
                         createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                         updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
