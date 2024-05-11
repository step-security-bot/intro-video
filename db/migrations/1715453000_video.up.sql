CREATE TABLE `videos` (
	`id` integer PRIMARY KEY NOT NULL,
  `url` text NOT NULL,
  `configuration_id` integer NOT NULL,
  FOREIGN KEY(`configuration_id`) REFERENCES `configurations`(`id`) ON UPDATE no action ON DELETE no action
);
