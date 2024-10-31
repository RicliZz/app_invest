create table if not exists startUp (
                         id SERIAL PRIMARY KEY,

    -- Информация о стартапе
                         title VARCHAR(100) NOT NULL,
                         topic VARCHAR(100),
                         idea TEXT,
                         strategy TEXT,
                         historyOfCreation TEXT,
                         status smallint,
                         stage smallint,

    -- Цель по финансированию, сколько нужно собрать
                         fundingGoal NUMERIC(15, 5),

    -- Предлагаемый процент инвестору
                         offeredPercent NUMERIC(3, 5),

    -- Информация о создателе
                         founderFullName VARCHAR(255),
                         founderEmail VARCHAR(255) UNIQUE,

    -- Ссылки
                         websiteLink VARCHAR(255),
                         vkLink VARCHAR(255),
                         telegramLink VARCHAR(255),
                         whatsUpLink VARCHAR(255),
                         instagramLink VARCHAR(255),

    -- Информация о дате создания и обновления
                         createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                         updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
