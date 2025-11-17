package items

type Item struct {
	ResourceName string `json:"resourceName"`
	TechDebtId   string `json:"techDebtId"` //todo need to add that these are required
}
