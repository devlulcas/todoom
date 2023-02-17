package db

type Pagination struct {
	Total        int64 `json:"total"`
	PerPage      int64 `json:"per_page"`
	Current      int64 `json:"current"`
	NextPage     int64 `json:"next_page"`
	PreviousPage int64 `json:"previous_page"`
}

type PaginationInput struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
