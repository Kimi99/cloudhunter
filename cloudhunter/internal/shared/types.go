package shared

type S3Node struct {
	Name     string
	IsFolder bool
	Children []*S3Node
}
