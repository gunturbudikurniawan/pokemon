package repository

import (
	"errors"
	"log"
	"pokemon/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Post_PokemonBattle(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		arg           models.BattleInput
	}

	var (
		testTable     []testCase
		expectedQuery = `
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
	)

	testTable = append(testTable, testCase{
		name:      "failed unexpected error",
		wantError: true,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("unexpected error"))
		},
		expectedError: errors.New("unexpected error"),
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		arg: models.BattleInput{
			Winner:    "ordinal",
			StartTime: time.Now(),
			EndTime:   time.Now(),
		},
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			if tc.mockQuery != nil {
				tc.mockQuery(mock)
			}
			repo := PokeRepo{
				db: db,
			}

			_, serr := repo.PostBattlePokemon(tc.arg)
			if tc.wantError {
				log.Print(tc.name)
				assert.EqualError(t, serr, tc.expectedError.Error())
			} else {
				log.Print(tc.name)
				assert.Nil(t, serr)
			}
		})
	}
}

func Test_Post_PokemonData(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		arg           models.Pokemon
	}

	var (
		testTable     []testCase
		expectedQuery = `			
		INSERT INTO
			pokemon(
				name,
				battle_id,
				scores
			)
		VALUES(
			$1, $2, $3
		)	`
	)

	testTable = append(testTable, testCase{
		name:      "failed unexpected error",
		wantError: true,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("unexpected error"))
		},
		expectedError: errors.New("unexpected error"),
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		arg: models.Pokemon{
			PokemonID: 1,
			Name:      "ahmad",
			BattleID:  2,
			Scores:    5,
		},
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			if tc.mockQuery != nil {
				tc.mockQuery(mock)
			}
			repo := PokeRepo{
				db: db,
			}

			serr := repo.PostPokemonData(tc.arg)
			if tc.wantError {
				log.Print(tc.name)
				assert.EqualError(t, serr, tc.expectedError.Error())
			} else {
				log.Print(tc.name)
				assert.Nil(t, serr)
			}
		})
	}
}

func Test_Get_PokemonScore(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		mockQuery      func(mock sqlmock.Sqlmock)
		expectedError  error
		expectedResult []models.DetailPlayers
	}

	var (
		testTable     []testCase
		expectedQuery = `
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

	testTable = append(testTable, testCase{
		name:      "failed unexpected error",
		wantError: true,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnError(errors.New("unexpected error"))
		},
		expectedError: errors.New("unexpected error"),
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnRows(sqlmock.NewRows([]string{"name", "scores"}).AddRow("pika", 2000).AddRow("chu", 1000))
		},
		expectedResult: []models.DetailPlayers{
			{
				Name:   "pika",
				Scores: 2000,
			},
			{
				Name:   "chu",
				Scores: 1000,
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			if tc.mockQuery != nil {
				tc.mockQuery(mock)
			}
			repo := PokeRepo{
				db: db,
			}

			res, serr := repo.GetPokemonScore()
			if tc.wantError {
				assert.EqualError(t, serr, tc.expectedError.Error())
			} else {
				assert.Nil(t, serr)
				assert.Equal(t, tc.expectedResult, res)
			}
		})
	}
}

func Test_Get_Player(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		mockQuery      func(mock sqlmock.Sqlmock)
		expectedError  error
		expectedResult []models.DetailPlayers
	}

	var (
		testTable     []testCase
		expectedQuery = `
		SELECT 
			p.name,
			p.scores
		FROM pokemon p 
		WHERE p.battle_id = $1
	`
	)

	testTable = append(testTable, testCase{
		name:      "failed unexpected error",
		wantError: true,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnError(errors.New("unexpected error"))
		},
		expectedError: errors.New("unexpected error"),
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnRows(sqlmock.NewRows([]string{"name", "scores"}).AddRow("Guntur", 200).AddRow("Kurniawan	", 100))
		},
		expectedResult: []models.DetailPlayers{
			{
				Name:   "Guntur",
				Scores: 200,
			},
			{
				Name:   "Kurniawan",
				Scores: 100,
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			if tc.mockQuery != nil {
				tc.mockQuery(mock)
			}
			repo := PokeRepo{
				db: db,
			}

			res, serr := repo.GetPlayer(1)
			if tc.wantError {
				log.Print(tc.name)
				assert.EqualError(t, serr, tc.expectedError.Error())
			} else {
				log.Print(tc.name)
				assert.Nil(t, serr)
				assert.Equal(t, tc.expectedResult, res)
			}
		})
	}
}

func Test_Get_Battle(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		mockQuery      func(mock sqlmock.Sqlmock)
		expectedError  error
		expectedResult []models.BattleResponse
	}

	var (
		testTable     []testCase
		expectedQuery = `
		SELECT 
			b.battle_id, 
			b.winner
		FROM battle b
	`
	)

	testTable = append(testTable, testCase{
		name:      "failed unexpected error",
		wantError: true,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnError(errors.New("unexpected error"))
		},
		expectedError: errors.New("unexpected error"),
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		mockQuery: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
				WillReturnRows(sqlmock.NewRows([]string{"battle_id", "winner"}).AddRow(1, "pikachu").AddRow(2, "pichu"))
		},
		expectedResult: []models.BattleResponse{
			{
				BattleID: 1,
				Winner:   "pikachu",
			},
			{
				BattleID: 2,
				Winner:   "pichu",
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			if tc.mockQuery != nil {
				tc.mockQuery(mock)
			}
			repo := PokeRepo{
				db: db,
			}

			res, serr := repo.GetBattle("2022-10-12", "2022-10-12")
			if tc.wantError {
				log.Print(tc.name)
				assert.EqualError(t, serr, tc.expectedError.Error())
			} else {
				log.Print(tc.name)
				assert.Nil(t, serr)
				assert.Equal(t, tc.expectedResult, res)
			}
		})
	}
}
