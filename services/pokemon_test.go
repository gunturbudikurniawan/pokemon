package services

import (
	"errors"
	"pokemon/models"
	postgres_mock "pokemon/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func Test_PokemonUsecase_GetAllPokemons(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResult   models.AllPokemon
		expectedError    error
		onGetAllPokemons func(mock *postgres_mock.MockPokemonRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:          "failed unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onGetAllPokemons: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetAllPokemons().Return(models.AllPokemon{}, errors.New("unexpected error")).AnyTimes()
		},
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onGetAllPokemons: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetAllPokemons().Return(models.AllPokemon{
				Count:    100,
				Next:     "",
				Previous: "",
				Results: []struct {
					Name string "json:\"name\""
					Url  string "json:\"url\""
				}{
					{
						Name: "pikachu",
						Url:  "wkwkwkkw",
					},
				},
			}, nil).Times(1)
		},
		expectedResult: models.AllPokemon{
			Count:    100,
			Next:     "",
			Previous: "",
			Results: []struct {
				Name string "json:\"name\""
				Url  string "json:\"url\""
			}{
				{
					Name: "pikachu",
					Url:  "wkwkwkkw",
				},
			},
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			pokeRepo := postgres_mock.NewMockPokemonRepo(mockCtrl)

			if testCase.onGetAllPokemons != nil {
				testCase.onGetAllPokemons(pokeRepo)
			}

			usecase := PokeUsecase{
				PokeRepository: pokeRepo,
			}

			data, serr := usecase.GetAllPokemons()

			if testCase.wantError {
				assert.EqualError(t, serr, testCase.expectedError.Error())
			} else {
				assert.Nil(t, serr)
				assert.Equal(t, testCase.expectedResult, data)
			}
		})
	}

}

func Test_PokemonUsecase_GetBattle(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult []models.BattleResponse
		expectedError  error
		onPokemonRepo  func(mock *postgres_mock.MockPokemonRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:          "failed unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onPokemonRepo: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetBattle(gomock.Any(), gomock.Any()).Return([]models.BattleResponse{}, errors.New("unexpected error")).AnyTimes()
			mock.EXPECT().GetPlayer(gomock.Any()).Return([]models.DetailPlayers{}, errors.New("unexpected error")).AnyTimes()
		},
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onPokemonRepo: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetBattle(gomock.Any(), gomock.Any()).Return([]models.BattleResponse{
				{
					BattleID: 1,
					Winner:   "pikachu",
				},
			}, nil).AnyTimes()
			mock.EXPECT().GetPlayer(gomock.Any()).Return([]models.DetailPlayers{
				{
					Name:   "pikachu",
					Scores: 68,
				},
			}, nil).AnyTimes()
		},
		expectedResult: []models.BattleResponse{
			{
				BattleID: 1,
				Winner:   "pikachu",
				Player: []models.DetailPlayers{
					{
						Name:   "pikachu",
						Scores: 68,
					},
				},
			},
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			pokeRepo := postgres_mock.NewMockPokemonRepo(mockCtrl)

			if testCase.onPokemonRepo != nil {
				testCase.onPokemonRepo(pokeRepo)
			}

			usecase := PokeUsecase{
				PokeRepository: pokeRepo,
			}

			data, serr := usecase.GetBattle("2022-07-02", "2022-07-02")

			if testCase.wantError {
				assert.EqualError(t, serr, testCase.expectedError.Error())
			} else {
				assert.Nil(t, serr)
				assert.Equal(t, testCase.expectedResult, data)
			}
		})
	}
}

func Test_PokemonUsecase_GetPokemonScore(t *testing.T) {
	type testCase struct {
		name              string
		wantError         bool
		expectedResult    []models.DetailPlayers
		expectedError     error
		onGetPokemonScore func(mock *postgres_mock.MockPokemonRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:          "failed unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onGetPokemonScore: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetPokemonScore().Return([]models.DetailPlayers{}, errors.New("unexpected error")).AnyTimes()
		},
	})

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onGetPokemonScore: func(mock *postgres_mock.MockPokemonRepo) {
			mock.EXPECT().GetPokemonScore().Return([]models.DetailPlayers{
				{
					Name:   "pikachu",
					Scores: 300,
				},
			}, nil).Times(1)
		},
		expectedResult: []models.DetailPlayers{
			{
				Name:   "pikachu",
				Scores: 300,
			},
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			pokeRepo := postgres_mock.NewMockPokemonRepo(mockCtrl)

			if testCase.onGetPokemonScore != nil {
				testCase.onGetPokemonScore(pokeRepo)
			}

			usecase := PokeUsecase{
				PokeRepository: pokeRepo,
			}

			data, serr := usecase.GetPokemonScore()

			if testCase.wantError {
				assert.EqualError(t, serr, testCase.expectedError.Error())
			} else {
				assert.Nil(t, serr)
				assert.Equal(t, testCase.expectedResult, data)
			}
		})
	}
}