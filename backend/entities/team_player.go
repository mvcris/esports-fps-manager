package entities

import "time"

type TeamPlayer struct {
	TeamID    string    `json:"teamID"`
	PlayerID  string    `json:"playerID"`
	IsCaptain bool      `json:"isCaptain"`
	JoinDate  time.Time `json:"joinDate"`
	LeaveDate time.Time `json:"leaveDate"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
