package client

import (
	"github.com/mxenabled/atrium-go/models"
)

func (c *Client) ListMembers(userGuid string) ([]*models.Member, error) {
	return nil, nil
}

func (c *Client) CreateMember(userGuid string, member *models.Member, credentials []*models.Credential) (*models.Member, error) {
	return nil, nil
}

func (c *Client) UpdateMember(userGuid string, member *models.Member, credentials []*models.Credential) (*models.Member, error) {
	return nil, nil
}

func (c *Client) AggregateMember(userGuid string, memberGuid string) (*models.Member, error) {
	return nil, nil
}

func (c *Client) ResumeMember(userGuid string, member *models.Member, challenges []*models.Challenge) (*models.Member, error) {
	return nil, nil
}
