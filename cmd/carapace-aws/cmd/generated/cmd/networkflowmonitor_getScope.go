package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getScopeCmd = &cobra.Command{
	Use:   "get-scope",
	Short: "Gets information about a scope, including the name, status, tags, and target details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getScopeCmd).Standalone()

	networkflowmonitor_getScopeCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
	networkflowmonitor_getScopeCmd.MarkFlagRequired("scope-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getScopeCmd)
}
