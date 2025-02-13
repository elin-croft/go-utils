package client

import "time"

type Option func(*Client)

func WithConnectionTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.TConfiguration.ConnectTimeout = t
	}
}
func WithSocketTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.TConfiguration.SocketTimeout = t
	}
}
