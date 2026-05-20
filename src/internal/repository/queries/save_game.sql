INSERT INTO games (id,player_x_id,player_o_id,status,winner_id,game_board,current_player,history,bot,bot_side,create_date)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
ON CONFLICT (id) DO UPDATE SET
		player_x_id = EXCLUDED.player_x_id,
		player_o_id = EXCLUDED.player_o_id,
		status = EXCLUDED.status,
		winner_id = EXCLUDED.winner_id,
		game_board = EXCLUDED.game_board,
		current_player = EXCLUDED.current_player,
		history = EXCLUDED.history,
		bot = EXCLUDED.bot,
		bot_side = EXCLUDED.bot_side,
		create_date = EXCLUDED.create_date