CREATE TABLE IF NOT EXISTS events(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    start_time timestamp(0) with time zone NOT NULL,
    end_time timestamp(0) with time zone NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    venue text NOT NULL,
    tags text[] NOT NULL,
    cover text NOT NULL,
    version integer NOT NULL DEFAULT 1
);