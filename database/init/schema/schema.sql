CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".permission ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_permission PRIMARY KEY ( id ),
	CONSTRAINT permission_name UNIQUE ( name ) 
 );

CREATE  TABLE "public"."role" ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_role PRIMARY KEY ( id ),
	CONSTRAINT role_name UNIQUE ( name ) 
 );

CREATE  TABLE "public".role_permission ( 
	id_role              integer  NOT NULL  ,
	id_permission        integer  NOT NULL  
 );

CREATE  TABLE "public".user_account ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	email                varchar(100)  NOT NULL  ,
	"password"           varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_user_account PRIMARY KEY ( id ),
	CONSTRAINT unq_user_email UNIQUE ( email ) 
 );

CREATE  TABLE "public".user_profile ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	user_id              integer  NOT NULL  ,
	age                  integer    ,
	phone                integer    ,
	district             varchar(100)    ,
	city                 varchar(100)    ,
	country              varchar(100)    ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_user_profile PRIMARY KEY ( id )
 );

CREATE  TABLE "public".user_role ( 
	id_user              integer  NOT NULL  ,
	id_role              integer  NOT NULL  
 );

ALTER TABLE "public".role_permission ADD CONSTRAINT fk_role_permission_permission FOREIGN KEY ( id_permission ) REFERENCES "public".permission( id ) ON DELETE CASCADE;

ALTER TABLE "public".role_permission ADD CONSTRAINT fk_role_permission_role FOREIGN KEY ( id_role ) REFERENCES "public"."role"( id ) ON DELETE CASCADE;

ALTER TABLE "public".user_profile ADD CONSTRAINT fk_user_profile_user_account FOREIGN KEY ( user_id ) REFERENCES "public".user_account( id ) ON DELETE CASCADE;

ALTER TABLE "public".user_role ADD CONSTRAINT fk_user_role_role FOREIGN KEY ( id_role ) REFERENCES "public"."role"( id ) ON DELETE CASCADE;

ALTER TABLE "public".user_role ADD CONSTRAINT fk_user_role_user_account FOREIGN KEY ( id_user ) REFERENCES "public".user_account( id ) ON DELETE CASCADE;
