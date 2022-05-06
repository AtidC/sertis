CREATE TABLE IF NOT EXISTS categories (
id      CHARACTER VARYING(5),
name    CHARACTER VARYING(20),
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
id      CHARACTER VARYING(5),
name    CHARACTER VARYING(20),
username   CHARACTER VARYING(20),
password   CHARACTER VARYING(20),
PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cards (
id                  CHARACTER VARYING(37),
category            CHARACTER VARYING(1),
title               CHARACTER VARYING(200),
status              CHARACTER VARYING(10),
content             CHARACTER VARYING(10000),
author              CHARACTER VARYING(5),
create_timestamp    TIMESTAMPTZ,
update_timestamp    TIMESTAMPTZ,
PRIMARY KEY (id)
);

INSERT INTO categories (id, name)
VALUES 
    ('1', 'Biology'),
    ('2', 'Finance'),
    ('3', 'Chemisty'),
    ('4', 'Engineering'),
    ('5', 'Health'),
    ('6', 'Society'),
    ('7', 'Space'),
    ('8', 'Art');

INSERT INTO users (id, name, username, password)
VALUES 
    ('1', 'Terrza Konecna', 'terrza', '1234'),
    ('2', 'Jana Novkova', 'jana', '1234'),
    ('3', 'Jakub Antlk', 'jakub', '1234');

INSERT INTO cards (id, category, title, status, content, author, create_timestamp, update_timestamp)
VALUES 
    ('706ef181-7418-442a-a512-c2b92dc981ed', '4', 'Rising seas could submerge Rio and Jakarta by 2100', 'publish', '
Aminath knows this all too well. As the environment and climate change minister for the Maldives, 
she is part of a community of politicians and scientists trying to work out how quickly sea levels will rise, 
if this can be slowed and what it means for us all. In some places, new ways of holding back the tide may buy us a few decades. 
Elsewhere, this won''t be possible. We are facing a disaster unfolding in slow motion. Responding effectively means a sea change in the way we think. …', 
'1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('54926313-8755-4749-8ec4-0b8756448d27', '1', 'How a soil microbe could rev up artificial photosynthesis', 'publish', '
Plants rely on a process called carbon fixation -- turning carbon dioxide from the air into carbon-rich biomolecules - for their very existence. 
That''s the whole point of photosynthesis, and a cornerstone of the vast interlocking system that cycles carbon through plants, 
animals, microbes and the atmosphere to sustain life on Earth.', 
'2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('fe51e2b2-e53d-44a5-a891-95cfa613f2f6', '5', 'Long-lasting healthy changes: Doable and worthwhile', 'publish', '
I''ve been a physician for 20 years now, and a strong proponent of lifestyle medicine for much of it. 
I know that it''s hard to make lasting, healthy lifestyle changes, even when people know what to do and have the means to do it. 
Yet many studies and my own clinical experience as a Lifestyle Medicine-certified physician have shown me a few approaches 
that can help make long-lasting healthy lifestyle changes happen.',
'3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
