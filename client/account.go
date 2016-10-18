package client

import (
	"github.com/mxenabled/atrium-go/models"
)

func (c *Client) ListAccounts(userGuid string) ([]*models.Account, error) {
	return nil, nil
}

func (c *Client) GetAccounts(userGuid, accountGuid string) (*models.Account, error) {
	return nil, nil
}

func (c *Client) ListAccountTransations(userGuid, accountGuid string) ([]*models.Transaction, error) {
	return nil, nil
}
