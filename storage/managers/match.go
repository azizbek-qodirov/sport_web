package managers

import (
	"database/sql"
	"fmt"
	"project/models"
)

type MatchManager struct {
	DB *sql.DB
}

func NewMatchManager(db *sql.DB) *MatchManager {
	return &MatchManager{DB: db}
}

func (m *MatchManager) GetMatches() (*models.MatchGARes, error) {
	query := `
		SELECT
			m.id,
			m.result1,
			m.result2,
			m.date,
			m.status,
			m.round,
			t1.id AS team1_id,
			t1.name AS team1_name,
			t1.country AS team1_country,
			t1.championship_id AS team1_championship_id,
			t2.id AS team2_id,
			t2.name AS team2_name,
			t2.country AS team2_country,
			t2.championship_id AS team2_championship_id,
			m.group_id
		FROM
			matches m
		JOIN
			teams t1 ON m.team1_id = t1.id
		JOIN
			teams t2 ON m.team2_id = t2.id
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %s", err.Error())
	}
	defer rows.Close()

	var matches models.MatchGARes

	for rows.Next() {
		var match models.Match
		var team1, team2 models.Team

		err := rows.Scan(
			&match.ID,
			&match.Result1,
			&match.Result2,
			&match.Date,
			&match.Status,
			&match.Round,
			&team1.ID,
			&team1.Name,
			&team1.Country,
			&team1.ChampionshipID,
			&team2.ID,
			&team2.Name,
			&team2.Country,
			&team2.ChampionshipID,
			&match.GroupID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %s", err.Error())
		}

		match.Team1 = team1
		match.Team2 = team2

		matches.Matches = append(matches.Matches, match)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %s", err.Error())
	}

	return &matches, nil
}

func (m *MatchManager) GetMatchByID(matchID string) (*models.Match, error) {
	query := `
		SELECT
			m.id,
			m.result1,
			m.result2,
			m.date,
			m.status,
			m.round,
			t1.id AS team1_id,
			t1.name AS team1_name,
			t1.country AS team1_country,
			t1.championship_id AS team1_championship_id,
			t2.id AS team2_id,
			t2.name AS team2_name,
			t2.country AS team2_country,
			t2.championship_id AS team2_championship_id,
			m.group_id,
			g.name AS group_name -- Add this line
		FROM
			matches m
		JOIN
			teams t1 ON m.team1_id = t1.id
		JOIN
			teams t2 ON m.team2_id = t2.id
		JOIN
			groups g ON m.group_id = g.id
		WHERE
			m.id = $1
	`

	row := m.DB.QueryRow(query, matchID)

	var match models.Match
	var team1, team2 models.Team

	err := row.Scan(
		&match.ID,
		&match.Result1,
		&match.Result2,
		&match.Date,
		&match.Status,
		&match.Round,
		&team1.ID,
		&team1.Name,
		&team1.Country,
		&team1.ChampionshipID,
		&team2.ID,
		&team2.Name,
		&team2.Country,
		&team2.ChampionshipID,
		&match.GroupID,
		&match.GroupName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning row: %s", err.Error())
	}

	match.Team1 = team1
	match.Team2 = team2

	return &match, nil
}
