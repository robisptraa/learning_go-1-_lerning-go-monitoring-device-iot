create table if not exists devices (
 id integer primary key AUTOINCREMENT,
 device_code text not null unique,
    device_name text not null,
    location text not null,
    status integer not null default 1,
    created_at timestamp default current_timestamp
);

create table if not exists sensor_logs (
    id integer primary key AUTOINCREMENT,
    device_id integer not null,
    temperature real not null,
    humidity real not null,
    created_at timestamp default current_timestamp,
    foreign key (device_id) references devices(id)
);



