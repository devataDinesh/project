package types

type Status interface {
	CheckStatus(websiteName string) (status bool, err error)
}
