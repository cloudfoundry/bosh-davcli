package cmd

import (
	"errors"
	"os"

	davclient "github.com/cloudfoundry/bosh-davcli/client"
)

type PutCmd struct {
	client davclient.Client
}

func newPutCmd(client davclient.Client) (cmd PutCmd) {
	cmd.client = client
	return
}

func (cmd PutCmd) Run(args []string) error {
	if len(args) != 2 {
		return errors.New("Incorrect usage, put needs local file and remote blob destination") //nolint:staticcheck
	}

	file, err := os.OpenFile(args[0], os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	return cmd.client.Put(args[1], file, fileInfo.Size())
}
