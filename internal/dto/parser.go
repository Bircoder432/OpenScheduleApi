package dto

type NewParserRequest struct {
	CollegeName string   `json:"collegeName"`
	CampusNames []string `json:"campusNames"`
}

type NewParserResponse struct {
	Token string `json:"token"`
}

type GetParserResponse struct {
	CollegeID uint `json:"collegeId"`
}
