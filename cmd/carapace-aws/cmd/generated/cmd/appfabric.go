package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabricCmd = &cobra.Command{
	Use:   "appfabric",
	Short: "Amazon Web Services AppFabric quickly connects software as a service (SaaS) applications across your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabricCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabricCmd).Standalone()

	})
	rootCmd.AddCommand(appfabricCmd)
}
