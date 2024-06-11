CREATE TABLE `videos` (
	`id` INTEGER PRIMARY KEY NOT NULL,
  `url` TEXT NOT NULL,
  `weight` INTEGER DEFAULT 1,
  `configuration_id` INTEGER NOT NULL,
  `instance_id` INTEGER NOT NULL,
  FOREIGN KEY(`configuration_id`) REFERENCES `configurations`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  FOREIGN KEY(`instance_id`) REFERENCES `instances`(`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
