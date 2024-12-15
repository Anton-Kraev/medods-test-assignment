package email

import "context"

type MockClient struct{}

func (MockClient) SendWarning(context.Context, string) error {
	return nil
}
