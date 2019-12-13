package dbx

import migrate "github.com/rubenv/sql-migrate"

func GetMigrations() []*migrate.Migration {
	return []*migrate.Migration{
		&migrate.Migration{
			Id: "1",
			Up: []string{`create table agency
						(
							id varchar(255),
							name varchar(255) null,
							url text null,
							timezone varchar(50) null,
							language varchar(150) null,
							phone varchar(50) null,
							constraint agency_pk
								primary key (id)
						);`},
			Down: []string{"DROP TABLE agency"},
		},
		&migrate.Migration{
			Id: "2",
			Up: []string{`create table route
						(
							id varchar(255),
							agency_id varchar(255) null,
							short_name varchar(50) null,
							long_name varchar(255) null,
							type int null,
							color varchar(50) null,
							text_color varchar(50) null,
							constraint route_pk
								primary key (id),
							constraint route_agency_id_fk
								foreign key (agency_id) references agency (id)
						);`},
			Down: []string{"DROP TABLE route"},
		},
		&migrate.Migration{
			Id: "3",
			Up: []string{`create table calendar
						(
							id int auto_increment,
							service_id varchar(255) not null,
							start_date varchar(50) null,
							end_date varchar(50) null,
							monday bool null,
							tuesday bool null,
							wednesday bool null,
							thursday bool null,
							friday bool null,
							saturday bool null,
							sunday bool null,
							constraint calendar_pk
								primary key (id)
						);`},
			Down: []string{"DROP TABLE calendar"},
		},
		&migrate.Migration{
			Id: "4",
			Up: []string{`create table calendar_date
						(
							id int auto_increment,
							service_id varchar(255) not null,
							date varchar(50) null,
							exception_type varchar(50) null,
							constraint calendar_date_pk
								primary key (id)
						);`},
			Down: []string{"DROP TABLE calendar_date"},
		},
		&migrate.Migration{
			Id: "5",
			Up: []string{`create table shape
						(
							id int auto_increment,
							shape_id varchar(255),
							latitude varchar(100) null,
							longitude varchar(100) null,
							sequence int null,
							distance_traveled float null,
							constraint shape_pk
								primary key (id)
						);`},
			Down: []string{"DROP TABLE shape"},
		},
	}
}
