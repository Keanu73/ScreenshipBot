-- Delete tables

DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS text_channel_subscriptions;
DROP TABLE IF EXISTS voice_channel_records;
DROP TABLE IF EXISTS voice_channel_categories;
DROP TABLE IF EXISTS settings;

DELETE TRIGGER settings_update_timestamp;
DELETE TRIGGER tcs_update_timestamp;
DELETE TRIGGER vcc_update_timestamp;
