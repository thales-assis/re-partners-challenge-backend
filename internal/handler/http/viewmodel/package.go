package viewmodel

type GetPackagesResponse struct {
	Sizes []int `json:"package_sizes" example:"[23, 31, 53]"`
}

type UpdateAllPackages struct {
	Sizes []int `json:"package_sizes" example:"[23, 31, 53]"`
}

func (vm UpdateAllPackages) Validate() error {

	validationErrors := NewValidationsError()

	for _, newPackageSize := range vm.Sizes {
		if newPackageSize <= 0 {
			validationErrors.Append("package_sizes", "must be a valid package site with value greater than 0")
		}
	}

	if validationErrors.HasError() {
		return validationErrors
	}

	return nil
}
