package models

import "time"

type MessageCReq struct {
	MatchID string    `json:"match_id" bson:"match_id"`
	Content string    `json:"content" bson:"content"`
	SentAt  time.Time `json:"sent_at" bson:"sent_at"`
}

type MessageGARes struct {
	Messages []*MessageCReq `json:"messages"`
}

type Team struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	ChampionshipID string `json:"championship_id"`
}

type Match struct {
	ID        string `json:"id"`
	Team1     Team   `json:"team1"`
	Team2     Team   `json:"team2"`
	GroupID   string `json:"group_id"`
	Result1   int    `json:"result1"`
	Result2   int    `json:"result2"`
	Date      string `json:"date"`
	Status    string `json:"status"`
	Round     int    `json:"round"`
	GroupName string `json:"group_name"`
}

type MatchGARes struct {
	Matches []Match `json:"matches"`
}

type MatchDetailsResponse struct {
	Match    *Match        `json:"match"`
	Messages *MessageGARes `json:"messages"`
	Group    string        `json:"group"`
}
