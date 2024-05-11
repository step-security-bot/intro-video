CREATE TABLE `configurations` (
	`id` integer PRIMARY KEY NOT NULL,
  `bubble_enabled` boolean DEFAULT false,
  `bubble_text_content` text,
  `cta_enabled` boolean DEFAULT false,
  `cta_text_content` text
);

