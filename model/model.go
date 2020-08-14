package model



type PodPostBodyUI struct {
	FirstName         string `json:"first_name"`
	LastName  				string `json:"last_name"`
	MobileNumber      string `json:"mobile_number"`
	PostalCode  			int 	 `json:"postal_code"`
	email    					string `json:"email"`
}
