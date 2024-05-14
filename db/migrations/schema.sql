CREATE TABLE schema_migrations (id VARCHAR(255) NOT NULL PRIMARY KEY);
CREATE TABLE `configurations` (
\t`id` integer PRIMARY KEY NOT NULL,
  `bubble_enabled` boolean DEFAULT false,
  `bubble_text_content` text,
  `bubble_type` text DEFAULT 'default',
  `cta_enabled` boolean DEFAULT false,
  `cta_text_content` text,
  `cta_type` text DEFAULT 'default'
);
CREATE TABLE `videos` (
\t`id` integer PRIMARY KEY NOT NULL,
  `url` text NOT NULL,
  `weight` integer DEFAULT 1,
  `configuration_id` integer NOT NULL,
  `instance_id` integer NOT NULL,
  FOREIGN KEY(`configuration_id`) REFERENCES `configurations`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  FOREIGN KEY(`instance_id`) REFERENCES `instances`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
CREATE TABLE `instances` (
\t`id` integer PRIMARY KEY NOT NULL
);