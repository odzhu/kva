package checkset

import (
	"context"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
)

//Etcdopened checks if etcd is publicly accessible
func (c *Check) Etcdopened() (*Check, error) {
	c.code = 201
	c.category = "security"
	c.description = "Etcd not protected"
	//fmt.Printf("%+v", c.checkset.Config.Host)
	if c.checkset.Config.Host != "" {
		u, err := url.Parse(c.checkset.Config.Host)
		h := strings.Split(u.Host, ":")
		host := h[0]
		//fmt.Println("Alive")
		cfg := client.Config{
			Endpoints: []string{"http://" + host + ":2379"},
			Transport: client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: time.Second,
		}
		cl, err := client.New(cfg)
		if err != nil {
			//fmt.Println("Alive")
			log.Fatal(err)
		}
		kapi := client.NewKeysAPI(cl)
		resp, err := kapi.Get(context.Background(), "/", nil)
		if err != nil {
			c.Result = false
		} else {
			//fmt.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
			c.resources = []string{resp.Node.Key, resp.Node.Value}
			c.Result = true
		}

	}
	return c, nil
}
