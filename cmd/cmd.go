package cmd

import (
	"github.com/sikalabs/signpost/cmd/root"
	_ "github.com/sikalabs/signpost/cmd/server"
	_ "github.com/sikalabs/signpost/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
