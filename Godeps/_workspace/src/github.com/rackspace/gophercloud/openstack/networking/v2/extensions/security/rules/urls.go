package rules

import "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"

const rootPath = "security-group-rules"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, id)
}
