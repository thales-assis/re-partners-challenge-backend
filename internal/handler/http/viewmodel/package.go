package viewmodel

type GetPacksResponse struct {
	Sizes []int `json:"pack_sizes" example:"[23, 31, 53]"`
}

type UpdateAllPacksRequest struct {
	Sizes []int `json:"pack_sizes" example:"[23, 31, 53]"`
}

func (vm UpdateAllPacksRequest) Validate() error {

	validationErrors := NewValidationsError()

	for _, newPackSize := range vm.Sizes {
		if newPackSize <= 0 {
			validationErrors.Append("pack_sizes", "must be a valid package site with value greater than 0")
		}
	}

	if validationErrors.HasError() {
		return validationErrors
	}

	return nil
}
