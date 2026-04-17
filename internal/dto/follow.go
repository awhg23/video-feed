package dto

type FollowUserItem struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type FollowListResponse struct {
	List []FollowUserItem `json:"list"`
}
