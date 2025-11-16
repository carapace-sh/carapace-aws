package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmakerCmd = &cobra.Command{
	Use:   "iottwinmaker",
	Short: "IoT TwinMaker is a service with which you can build operational digital twins of physical systems.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmakerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmakerCmd).Standalone()

	})
	rootCmd.AddCommand(iottwinmakerCmd)
}
