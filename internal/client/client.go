package client

import (
	"context"
	"fmt"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/vladimirok5959/golang-ip2location/internal/consts"
)

type Client struct {
	ctx  context.Context
	base *ip2location.DB
}

type Result struct {
	City         string
	CountryLong  string
	CountryShort string
	Region       string
}

func New(ctx context.Context, shutdown context.CancelFunc) (*Client, error) {
	f, err := consts.DataPathFile("IP2LOCATION-LITE-DB3.BIN")
	if err != nil {
		return nil, err
	}

	b, err := ip2location.OpenDB(f)
	if err != nil {
		return nil, err
	}

	c := Client{
		ctx:  ctx,
		base: b,
	}

	return &c, nil
}

func (c *Client) IP2Location(ctx context.Context, ip string) (*Result, error) {
	if c.base == nil {
		return nil, fmt.Errorf("database is not opened")
	}

	r, err := c.base.Get_all(ip)

	return &Result{
		City:         r.City,
		CountryLong:  r.Country_long,
		CountryShort: r.Country_short,
		Region:       r.Region,
	}, err
}

func (c *Client) ReloadDatabase(ctx context.Context) error {
	f, err := consts.DataPathFile("IP2LOCATION-LITE-DB3.BIN")
	if err != nil {
		return err
	}

	b, err := ip2location.OpenDB(f)
	if err != nil {
		return err
	}

	c.base.Close()
	c.base = b

	return nil
}

func (c *Client) Shutdown(ctx context.Context) error {
	c.base.Close()
	return nil
}
