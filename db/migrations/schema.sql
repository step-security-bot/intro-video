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
  `configuration_id` integer NOT NULL,
  FOREIGN KEY(`configuration_id`) REFERENCES `configurations`(`id`) ON UPDATE no action ON DELETE no action
);
CREATE TABLE `instances` (
\t`id` integer PRIMARY KEY NOT NULL,
  `video1_id` integer,
  `video2_id` integer,
  `video3_id` integer,
  `video4_id` integer,
  `video5_id` integer,
  `video1_weight` integer DEFAULT 1,
  `video2_weight` integer DEFAULT 1,
  `video3_weight` integer DEFAULT 1,
  `video4_weight` integer DEFAULT 1,
  `video5_weight` integer DEFAULT 1,
  FOREIGN KEY(`video1_id`) REFERENCES `videos`(`id`) ON UPDATE no action ON DELETE no action,
  FOREIGN KEY(`video2_id`) REFERENCES `videos`(`id`) ON UPDATE no action ON DELETE no action,
  FOREIGN KEY(`video3_id`) REFERENCES `videos`(`id`) ON UPDATE no action ON DELETE no action,
  FOREIGN KEY(`video4_id`) REFERENCES `videos`(`id`) ON UPDATE no action ON DELETE no action,
  FOREIGN KEY(`video5_id`) REFERENCES `videos`(`id`) ON UPDATE no action ON DELETE no action
);
