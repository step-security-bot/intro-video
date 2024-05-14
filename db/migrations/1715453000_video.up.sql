CREATE TABLE `videos` (
	`id` integer PRIMARY KEY NOT NULL,
  `url` text NOT NULL,
  `weight` integer DEFAULT 1,
  `configuration_id` integer NOT NULL,
  `instance_id` integer NOT NULL,
  FOREIGN KEY(`configuration_id`) REFERENCES `configurations`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  FOREIGN KEY(`instance_id`) REFERENCES `instances`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
