CREATE TABLE IF NOT EXISTS voices(
    user_id uuid NOT NULL,
    petition_id uuid NOT NULL,
    FOREIGN KEY(petition_id) REFERENCES petitions(id)
)