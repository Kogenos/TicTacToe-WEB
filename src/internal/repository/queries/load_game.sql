SELECT id,player_x_id,player_o_id,status,winner_id,game_board,current_player,history,bot,bot_side,create_date
FROM games
WHERE id = $1
	