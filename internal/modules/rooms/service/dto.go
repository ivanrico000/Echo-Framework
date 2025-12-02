package application

type RoomCreateRequest struct {
	Number      int    `json:"number"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SingleBeds  int    `json:"single_beds"`
	DoubleBeds  int    `json:"double_beds"`
}

type RoomUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SingleBeds  *int    `json:"single_beds,omitempty"`
	DoubleBeds  *int    `json:"double_beds,omitempty"`
	Number      *int    `json:"number,omitempty"`
}

type RoomResponse struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SingleBeds  int    `json:"single_beds"`
	DoubleBeds  int    `json:"double_beds"`
}
