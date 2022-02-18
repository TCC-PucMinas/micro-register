package helpers

type Paginate struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

func (p *Paginate) PaginateMounted() {
	switch p.Page {
	case 0:
		p.Page = 0
	case 1:
		p.Page = 0
	default:
		p.Page = (p.Page * 10) - 10
	}
}
