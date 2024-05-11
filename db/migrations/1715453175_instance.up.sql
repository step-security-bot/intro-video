CREATE TABLE `instances` (
	`id` integer PRIMARY KEY NOT NULL,
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

