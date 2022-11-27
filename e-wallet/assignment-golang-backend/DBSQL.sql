create sequence seq_wallet
start 100000
increment 1;

create table wallets(
	wallet_id INT primary key default nextval('seq_wallet'), 
 	balance int not NULL, 
	created_at TIMESTAMP DEFAULT current_timestamp,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
)


create table users(
	user_id SERIAL PRIMARY KEY,
	email VARCHAR(255) not NULL, 
	wallet_id  int, 	
	password text not NULL, 
	created_at TIMESTAMP DEFAULT current_timestamp,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	foreign key (wallet_id) references wallets(wallet_id) on update cascade on delete cascade
)

create table transactions (
	transaction_id SERIAL PRIMARY KEY,
 	transaction_type  VARCHAR(255) not NULL, 
 	sender_id int not NULL,
 	receiver_id int not NULL,
 	amount int not NULL,
 	description  VARCHAR(255) not NULL, 
 	source_fund_id int not NULL,	
	created_at TIMESTAMP DEFAULT current_timestamp,
	updated_at TIMESTAMP DEFAULT current_timestamp,
	deleted_at TIMESTAMP,
	foreign key (sender_id) references users(user_id) on update cascade on delete cascade,
	foreign key (receiver_id) references users(user_id) on update cascade on delete cascade
)

INSERT INTO public.users (email,wallet_id,"password",created_at,updated_at,deleted_at) VALUES
	 ('user1@gmail.com',100001,'$2a$04$xCiqG0ytN35ioZUoncAUZOAqnQjwsUKnJGqlJ/SkLfIKwxdwCjY..','2022-10-13 16:19:10.291897','2022-10-13 16:19:10.291897',NULL),
	 ('user2@gmail.com',100002,'$2a$04$oYkdmfLtSkxE5muVL66hlu7LPl246U0rhno1z3mmS9AHiJCe47bKS','2022-10-13 16:19:16.920711','2022-10-13 16:19:16.920711',NULL),
	 ('user3@gmail.com',100003,'$2a$04$1Hv3jvB5gW.IwjP9m4W.S.y/HsC1z7I59hNt4vfjyN.zkd1r1uwmK','2022-10-13 16:19:23.083911','2022-10-13 16:19:23.083911',NULL),
	 ('user4@gmail.com',100004,'$2a$04$9sFqkOewMoJvg2N7vZh.Y.dFmf8yLvyVuEVgu7vkYiD6dRdSDWscW','2022-10-13 16:19:28.661316','2022-10-13 16:19:28.661316',NULL),
	 ('user5@gmail.com',100005,'$2a$04$dVMgE2nTUrtsAuoZs1UlGuba0fddMDPngwGN0BBoxf.iKy9fkOCB6','2022-10-13 16:19:34.150708','2022-10-13 16:19:34.150708',NULL);

INSERT INTO public.transactions (transaction_type,sender_id,receiver_id,amount,description,source_fund_id,created_at,updated_at,deleted_at) VALUES
	 ('TOP UP',1,1,500000,'Top Up from Credit Card',2,'2022-10-13 16:22:51.254072','2022-10-13 16:22:51.254072',NULL),
	 ('TRANSFER',1,2,10000,'desc1',0,'2022-10-13 16:22:57.640009','2022-10-13 16:22:57.640009',NULL),
	 ('TRANSFER',1,3,10000,'desc1',0,'2022-10-13 16:23:02.437065','2022-10-13 16:23:02.437065',NULL),
	 ('TRANSFER',1,4,10000,'desc1',0,'2022-10-13 16:23:06.150937','2022-10-13 16:23:06.150937',NULL),
	 ('TRANSFER',1,5,10000,'desc5',0,'2022-10-13 16:23:13.713875','2022-10-13 16:23:13.713875',NULL),
	 ('TOP UP',1,1,500000,'Top Up from Credit Card',2,'2022-10-13 16:24:19.876474','2022-10-13 16:24:19.876474',NULL),
	 ('TRANSFER',1,1,10000,'desc1',0,'2022-10-13 16:24:34.035099','2022-10-13 16:24:34.035099',NULL),
	 ('TRANSFER',1,2,10000,'desc1',0,'2022-10-13 16:24:38.287762','2022-10-13 16:24:38.287762',NULL),
	 ('TRANSFER',1,3,10000,'desc1',0,'2022-10-13 16:24:42.548933','2022-10-13 16:24:42.548933',NULL),
	 ('TRANSFER',1,4,10000,'desc1',0,'2022-10-13 16:24:45.725762','2022-10-13 16:24:45.725762',NULL),
	 ('TOP UP',2,2,500000,'Top Up from Credit Card',2,'2022-10-13 16:25:38.286183','2022-10-13 16:25:38.286183',NULL),
	 ('TRANSFER',2,1,10000,'desc1',0,'2022-10-13 16:25:56.327092','2022-10-13 16:25:56.327092',NULL),
	 ('TRANSFER',2,2,30000,'desc1',0,'2022-10-13 16:26:05.759081','2022-10-13 16:26:05.759081',NULL),
	 ('TRANSFER',2,3,40000,'desc1',0,'2022-10-13 16:26:10.565001','2022-10-13 16:26:10.565001',NULL),
	 ('TOP UP',3,3,500000,'Top Up from Credit Card',2,'2022-10-13 16:28:39.13711','2022-10-13 16:28:39.13711',NULL),
	 ('TRANSFER',3,1,40000,'desc1',0,'2022-10-13 16:30:51.600007','2022-10-13 16:30:51.600007',NULL),
	 ('TRANSFER',3,2,30000,'desc1',0,'2022-10-13 16:30:59.253548','2022-10-13 16:30:59.253548',NULL),
	 ('TRANSFER',3,5,10000,'desc1',0,'2022-10-13 16:32:57.983398','2022-10-13 16:32:57.983398',NULL),
	 ('TRANSFER',3,1,12000,'desc1',0,'2022-10-13 16:33:44.991053','2022-10-13 16:33:44.991053',NULL),
	 ('TRANSFER',3,3,13000,'desc1',0,'2022-10-13 16:33:52.581442','2022-10-13 16:33:52.581442',NULL);

INSERT INTO public.wallets (balance,created_at,updated_at,deleted_at) VALUES
	 (20000,'2022-10-13 16:19:28.660062','2022-10-13 16:19:28.660233',NULL),
	 (500000,'2022-10-13 16:19:16.919066','2022-10-13 16:19:16.919289',NULL),
	 (20000,'2022-10-13 16:19:34.149355','2022-10-13 16:19:34.149534',NULL),
	 (992000,'2022-10-13 16:19:10.288219','2022-10-13 16:19:10.288737',NULL),
	 (468000,'2022-10-13 16:19:23.082469','2022-10-13 16:19:23.082689',NULL);
