package root

import (
	"github.com/sikalabs/signpost/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "signpost",
	Short: "signpost, " + version.Version,
}
