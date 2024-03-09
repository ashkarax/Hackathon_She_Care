package requestmodel

type PostData struct {
	Title   string `json:"title" validate:"required,gte=4,lte=50"`
	Content string `json:"content" validate:"required,gte=30,lte=500"`
}
