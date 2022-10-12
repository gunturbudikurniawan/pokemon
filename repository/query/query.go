package query

const (
	PostBattle = `			
		INSERT INTO
			battle(
				winner,
				start_time,
				end_time
			)
		VALUES(
			$1, $2, $3
		)
		RETURNING battle_id;
	`

	PostPokemon = `			
		INSERT INTO
			pokemon(
				name,
				battle_id,
				scores
			)
		VALUES(
			$1, $2, $3
		)
	`

	GetBattle = `
		SELECT 
			b.battle_id, 
			b.winner
		FROM battle b 
	`

	GetPlayers = `
		SELECT 
			p.name,
			p.scores
		FROM pokemon p 
		WHERE p.battle_id = $1
	`

	GetPokemonScore = `
		SELECT
 			name,
		SUM (scores) AS scores
		FROM
			pokemon
		GROUP BY
			name
		ORDER BY scores desc;	
	`
)
