-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Etc/GMT";

-- store ADMIN_ROLE_IDS in .env or something
CREATE TYPE setting_name AS ENUM ('action_logging_channel_id', 'auto_kick_duration', 'character_limit', 'welcome_message');
CREATE TYPE voice_channel_category AS ENUM ('meditate', 'journal', 'gratitude', 'exercise', 'storytelling', 'deep-work');

-- We need to store different categories depending on the voice channels for action logging
-- E.g. user joins/leaves channel, Golang code writes to database. If channel ID is part of the 'meditation' category (modified dynamically), then, when message is posted to #action-logging, we can say "user meditated"!

CREATE TABLE settings (
    name setting_name NOT NULL PRIMARY KEY,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL DEFAULT NOW ()
);

CREATE TABLE text_channel_subscriptions (
    role_id BIGINT NOT NULL PRIMARY KEY,
    channel_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL DEFAULT NOW ()
);

CREATE TABLE voice_channel_records (
    user_id BIGINT NOT NULL PRIMARY KEY,
    channel_id BIGINT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    recorded_at TIMESTAMP DEFAULT NOW ()
);

CREATE TABLE voice_channel_categories (
    channel_id BIGINT NOT NULL PRIMARY KEY,
    category voice_channel_category NOT NULL,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL DEFAULT NOW ()
);

/*CREATE TABLE real_talk_voice_channels (
    channel_id BIGINT NOT NULL PRIMARY KEY,
    max_people INTEGER NOT NULL,
    locked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL DEFAULT NOW ()
);*/

-- Add indexes
CREATE INDEX vcr_channel_id ON voice_channel_records (channel_id);

-- thanks Theo F https://stackoverflow.com/a/68503819

--create function trigger to change a timestamp value upon an update
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--create a trigger to execute the function
CREATE TRIGGER settings_update_timestamp
    BEFORE UPDATE ON settings
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER tcs_update_timestamp
    BEFORE UPDATE ON text_channel_subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER vcc_update_timestamp
    BEFORE UPDATE ON voice_channel_categories
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

