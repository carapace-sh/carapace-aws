package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaverCmd = &cobra.Command{
	Use:   "simspaceweaver",
	Short: "SimSpace Weaver (SimSpace Weaver) is a service that you can use to build and run large-scale spatial simulations in the Amazon Web Services Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaverCmd).Standalone()

	})
	rootCmd.AddCommand(simspaceweaverCmd)
}
