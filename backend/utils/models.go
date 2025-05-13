package utils

// Comment represents a user comment on a post
type Comment struct {
	Id         int    `json:"id"`
	PostId     int    `json:"postId"`
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	Content    string `json:"content"`
	ImagePath  string `json:"imagePath"`
	Date       int64  `json:"date"`
}

// Reaction represents a user reaction to a post (like, love, etc.)
type Reaction struct {
	Id           int    `json:"id"`
	PostId       int    `json:"postId"`
	UserId       int    `json:"userId"`
	ReactionType string `json:"reactionType"`
	Date         int64  `json:"date"`
}

// ReactionCount represents the count of each reaction type for a post
type ReactionCount struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
