CREATE TABLE `videos` (
	`id` integer PRIMARY KEY NOT NULL,
  `script_id` integer NOT NULL,
  `stylesheet_id` integer NOT NULL,
  FOREIGN KEY(`script_id`) REFERENCES `scripts`(`id`) ON UPDATE no action ON DELETE no action,
  FOREIGN KEY(`stylesheet_id`) REFERENCES `scripts`(`id`) ON UPDATE no action ON DELETE no action
);
