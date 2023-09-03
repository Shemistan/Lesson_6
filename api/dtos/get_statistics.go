package dtos

type GetStatisticsRequest struct {
}

type GetStatisticsResponse struct {
	DeletedUsersCount int
	UpdateCount       int
	GetUserCount      int
	GetUsersCount     int
}
