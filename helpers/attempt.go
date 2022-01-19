package helpers

const (
	integrationRetry = 5
)

func EqualRetry(err error, retry int) bool {
	return err != nil && retry < integrationRetry
}
