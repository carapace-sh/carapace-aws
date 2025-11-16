package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_updateScopeCmd = &cobra.Command{
	Use:   "update-scope",
	Short: "Update a scope to add or remove resources that you want to be available for Network Flow Monitor to generate metrics for, when you have active agents on those resources sending metrics reports to the Network Flow Monitor backend.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_updateScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_updateScopeCmd).Standalone()

		networkflowmonitor_updateScopeCmd.Flags().String("resources-to-add", "", "A list of resources to add to a scope.")
		networkflowmonitor_updateScopeCmd.Flags().String("resources-to-delete", "", "A list of resources to delete from a scope.")
		networkflowmonitor_updateScopeCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_updateScopeCmd.MarkFlagRequired("scope-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_updateScopeCmd)
}
