package adapter

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("client inserts Lightning connector into computer.")
	com.InsertIntoLightingPort()
}
