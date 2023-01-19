package server

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sikalabs/signpost/cmd/root"
	"github.com/sikalabs/signpost/pkg/server"
	"github.com/spf13/cobra"
)

var FlagConfigFile string

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "run server",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		config := server.Config{}

		configFileBytes, err := ioutil.ReadFile(FlagConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(configFileBytes, &config)
		if err != nil {
			log.Fatal(err)
		}

		err = server.Server(config)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&FlagConfigFile,
		"config",
		"c",
		"",
		"Config file",
	)
	Cmd.MarkPersistentFlagRequired("config")
}
