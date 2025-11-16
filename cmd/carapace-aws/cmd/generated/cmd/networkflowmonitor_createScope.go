package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_createScopeCmd = &cobra.Command{
	Use:   "create-scope",
	Short: "In Network Flow Monitor, you specify a scope for the service to generate metrics for.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_createScopeCmd).Standalone()

	networkflowmonitor_createScopeCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters that you specify to make an idempotent API request.")
	networkflowmonitor_createScopeCmd.Flags().String("tags", "", "The tags for a scope.")
	networkflowmonitor_createScopeCmd.Flags().String("targets", "", "The targets to define the scope to be monitored.")
	networkflowmonitor_createScopeCmd.MarkFlagRequired("targets")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_createScopeCmd)
}
