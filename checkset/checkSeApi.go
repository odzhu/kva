package checkset

import "strings"

//Apiinsecure checks if etcd is publicly accessible
func (c *Check) Apiinsecure() (*Check, error) {
	c.code = 202
	c.category = "security"
	c.description = "Api server public availability check"

	h := strings.Split(c.checkset.Config.Host, ":")
	//TLS enabled ?
	if h[0] != "https" {
		c.Result = true
	} else {
		c.Result = false
	}
	//fmt.Printf("%+v", c.checkset.Config.Host)
	c.resources = append(c.resources, c.checkset.Config.Host)
	return c, nil
}
