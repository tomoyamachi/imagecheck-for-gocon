package pkg

import (
	"fmt"
	"log"

	"github.com/tomoyamachi/imagecheck-for-gocon/pkg/nginx"

	"github.com/genuinetools/reg/registry"
	deckodertypes "github.com/goodwithtech/deckoder/types"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) (err error) {
	args := c.Args()
	if len(args) == 0 {
		log.Println("requires at least 1 argument or --input option.")
		cli.ShowAppHelpAndExit(c, 1)
		return
	}

	imageName := args[0]
	// Check whether 'latest' tag is used
	_, err = registry.ParseImage(imageName)
	if err != nil {
		return fmt.Errorf("invalid image: %w", err)
	}

	// set docker option
	dockerOption := deckodertypes.DockerOption{
		Timeout:  c.Duration("timeout"),
		AuthURL:  c.String("authurl"),
		UserName: c.String("username"),
		Password: c.String("password"),
		Insecure: c.BoolT("insecure"),
		NonSSL:   c.BoolT("nonssl"),
	}
	log.Println("Start assessments...")

	err = nginx.ScanImage(imageName, dockerOption)
	if err != nil {
		return err
	}

	// Store ignore checkpoint code
	//output := os.Stdout
	return nil
}
