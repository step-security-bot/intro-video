CREATE TABLE `instances` (
	`id` INTEGER PRIMARY KEY NOT NULL,
  `uuid` BLOB UNIQUE NOT NULL
);

CREATE INDEX instances_uuid_idx ON instances(uuid);
