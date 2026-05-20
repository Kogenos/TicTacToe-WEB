DROP TABLE if EXISTS games;

create TABLE if not exists games (
    id UUID PRIMARY KEY,
    player_x_id UUID NOT NULL,
    player_o_id UUID,
    status TEXT NOT NULL,
    winner_id UUID,
    game_board JSONB NOT NULL,
    current_player int NOT NULL,
    history JSONB NOT NULL,
    bot BOOLEAN NOT NULL,
    bot_side INT NOT NULL,
    create_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);