package resources

type Resource struct {
	ResourceId    string `json:"resource_id"`
	ResourceName  string `json:"resource_name"`
	Os            string `json:"os"`
	ApplicationId string `json:"application_id"`
	ProductId     string `json:"product_id"`
}
