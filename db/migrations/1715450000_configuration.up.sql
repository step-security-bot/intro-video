CREATE TABLE `configurations` (
	`id` INTEGER PRIMARY KEY NOT NULL,
  `theme` TEXT DEFAULT 'default',
  `bubble_enabled` BOOLEAN DEFAULT false,
  `bubble_text_content` TEXT,
  `cta_enabled` BOOLEAN DEFAULT false,
  `cta_text_content` TEXT
);

