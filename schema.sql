CREATE TABLE users (
    id integer primary key autoincrement,
    username text not null,
    password text not null,
    created_at timestamp not null default current_timestamp
);

CREATE TABLE songs (
    id integer primary key autoincrement,
    created_by integer not null,
    created_at timestamp not null default current_timestamp,
    list_id integer not null,
    link text not null,
    comment text,
    foreign key(created_by) references users(id) on delete cascade,
    foreign key(list_id) references lists(id) on delete cascade
);

CREATE TABLE users_lists (
    id integer primary key autoincrement,
    user_id integer not null,
    list_id integer not null,
    foreign key(user_id) references users(id),
    foreign key(list_id) references song_lists(id)
);

CREATE TABLE list_invites (
    id integer primary key autoincrement,
    created_by integer not null,
    code text not null,
    list_id integer not null,
    invite_uses integer not null default 0,
    created_at timestamp not null default current_timestamp,
    foreign key(created_at) references users(id),
    foreign key(list_id) references song_lists(id)
);

CREATE TABLE lists (
    id integer primary key autoincrement,
    created_at timestamp not null default current_timestamp,
    created_by integer not null,
    hidden integer not null default 0,
    name text not null,
    description text,
    foreign key(created_by) references users(id) on delete cascade
);

CREATE TABLE song_comments (
    id integer primary key autoincrement,
    created_by integer not null,
    song_id integer not null,
    comment text,
    foreign key(created_by) references users(id) on delete cascade,
    foreign key(song_id) references songs(id) on delete cascade
);

CREATE TABLE song_votes (
    user_id integer not null,
    song_id integer not null,
    value integer not null,
    foreign key(user_id) references users(id) on delete cascade,
    foreign key(song_id) references songs(id) on delete cascade,
    CHECK (value = 1 OR value = -1)
);
