package models

type ResourceInfo struct {
	ResourceType string `json:"resourceType"` //资源类别, 带宽, 算力等
	Percentage   int    `json:"percentage"`   //资源使用百分比
}
