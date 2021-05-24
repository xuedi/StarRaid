DROP TABLE IF EXISTS `objects`;
CREATE TABLE `objects`
(
    `id`   varchar(36) NOT NULL,
    `name` varchar(30) NOT NULL,
    `x`    bigint(20) NOT NULL,
    `y`    bigint(20) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO objects (id, name, x, y)
VALUES ('8df5204b-a454-4c52-96bc-71ce03b5c38e', 'Ship 1', -200, -200),
       ('640ee436-08c3-4a4c-a9a9-1b58008a4d6b', 'Ship 2', 200, 200),
       ('7669521b-98ba-4258-868e-512497850f7f', 'Ship 3', 600, 200),
       ('9276357c-a45d-4ae5-b8c5-8bb9313d98ce', 'Pirate 1', 1000, 100),
       ('44f7fea7-3340-4f25-9338-4abe6e7ad72e', 'Pirate 2', 1000, 200),
       ('c89c266a-2010-4d56-a6dd-98961ad40126', 'Pirate 3', 1000, 300),
       ('358f052f-1d32-4a4d-95ec-7cbc46e04944', 'Pirate 4', 1000, 400),
       ('190d9202-1473-4a15-9889-b90bf0c6aa99', 'Pirate 5', 50000, 100),
       ('fbf44c0c-3d07-486e-8bc5-25035c7d098c', 'Pirate 6', 50000, 200),
       ('b09f2996-ea7f-402e-b173-7895764a961c', 'Pirate 7', 50000, 300),
       ('3bd51f42-f0b9-4b43-ab02-5f1792b90f54', 'Pirate 8', 50000, 400),
       ('65e74c19-9b8d-4b83-b5c6-b4a1784a6640', 'Station', 0, 0);



DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`        varchar(36) NOT NULL,
    `name`      varchar(30) NOT NULL,
    `object_id` varchar(36) NULL,
    `password`  varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO user (id, name, object_id, password)
VALUES ('9385c0e9-baf6-4e2e-bb7b-a5b6b4a0dd49', 'NPC', null, '1234'),
       ('213872bd-7c32-43d3-9933-076681362576', 'Player1', '640ee436-08c3-4a4c-a9a9-1b58008a4d6b', '1234'),
       ('7f361c30-2e37-47db-858b-45981350cb60', 'Player2', '7669521b-98ba-4258-868e-512497850f7f', '1234'),
       ('6dddea44-9a70-4843-94c7-b772e0e1ae7c', 'xuedi', '8df5204b-a454-4c52-96bc-71ce03b5c38e', '1234');

