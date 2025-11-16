package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_deleteScopeCmd = &cobra.Command{
	Use:   "delete-scope",
	Short: "Deletes a scope that has been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_deleteScopeCmd).Standalone()

	networkflowmonitor_deleteScopeCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
	networkflowmonitor_deleteScopeCmd.MarkFlagRequired("scope-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_deleteScopeCmd)
}
