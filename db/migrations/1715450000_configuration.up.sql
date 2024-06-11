CREATE TABLE `configurations` (
	`id` INTEGER PRIMARY KEY NOT NULL,
  `bubble_enabled` BOOLEAN DEFAULT false,
  `bubble_text_content` TEXT,
  `bubble_type` TEXT DEFAULT 'default',
  `cta_enabled` BOOLEAN DEFAULT false,
  `cta_text_content` TEXT,
  `cta_type` TEXT DEFAULT 'default'
);

