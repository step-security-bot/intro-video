CREATE TABLE `instances` (
	`id` INTEGER PRIMARY KEY NOT NULL,
  `external_id` BLOB UNIQUE NOT NULL
);

CREATE INDEX instances_external_id_idx ON instances(external_id);
