WITH user_games AS (

    SELECT
        player_x_id AS user_id,
        CASE
            WHEN winner_id = player_x_id THEN 1
            ELSE 0
        END AS win,
        CASE
            WHEN status = 'finished' AND winner_id IS NOT NULL AND winner_id != player_x_id THEN 1
            ELSE 0
            END AS loss,
        CASE 
            WHEN status = 'finished' AND winner_id IS NULL THEN 1
            ELSE 0
        END AS draw
    FROM games
    WHERE status = 'finished'

    UNION ALL

    SELECT
        player_o_id AS user_id,
        CASE
            WHEN winner_id = player_o_id THEN 1
            ELSE 0
        END AS win,
        CASE
            WHEN status = 'finished' AND winner_id IS NOT NULL AND winner_id != player_o_id THEN 1
            ELSE 0
            END AS loss,
        CASE 
            WHEN status = 'finished' AND winner_id IS NULL THEN 1
            ELSE 0
        END AS draw
    FROM games
    WHERE status = 'finished' AND player_o_id IS NOT NULL
),
aggregated AS (

    SELECT 
        user_id,
        SUM(win) AS wins,
        SUM(loss) AS losses,
        SUM(draw) AS draws,
        COUNT(*) AS total_games
    FROM user_games
    GROUP BY user_id
    HAVING COUNT(*) > 0
)
SELECT
    u.id,
    u.login,
    COALESCE(a.wins,0) AS wins,
    COALESCE(a.total_games,0) AS total_games,
    CASE
        WHEN COALESCE(a.total_games,0) = 0 THEN 0.0
        ELSE CAST(a.wins AS FLOAT) / a.total_games
    END AS win_ratio
    FROM users u
    LEFT JOIN aggregated a ON u.id = a.user_id
    ORDER BY win_ratio DESC, wins DESC
    LIMIT $1



