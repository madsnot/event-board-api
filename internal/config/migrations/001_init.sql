CREATE TABLE users(
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    username TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    avatar_url TEXT NOT NULL DEFAULT '',
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    middlename TEXT NOT NULL,
    gender TEXT NOT NULL,
    birthdate TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE sessions(
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE events(
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    author_id UUID NOT NULL,
    title TEXT NOT NULL,
    type INT NOT NULL,
    theme TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    genders JSONB NOT NULL,
    age INT NOT NULL DEFAULT 0,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    closed_at TIMESTAMP
);
