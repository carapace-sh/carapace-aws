package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_associateConfigurationItemsToApplicationCmd = &cobra.Command{
	Use:   "associate-configuration-items-to-application",
	Short: "Associates one or more configuration items with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_associateConfigurationItemsToApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_associateConfigurationItemsToApplicationCmd).Standalone()

		discovery_associateConfigurationItemsToApplicationCmd.Flags().String("application-configuration-id", "", "The configuration ID of an application with which items are to be associated.")
		discovery_associateConfigurationItemsToApplicationCmd.Flags().String("configuration-ids", "", "The ID of each configuration item to be associated with an application.")
		discovery_associateConfigurationItemsToApplicationCmd.MarkFlagRequired("application-configuration-id")
		discovery_associateConfigurationItemsToApplicationCmd.MarkFlagRequired("configuration-ids")
	})
	discoveryCmd.AddCommand(discovery_associateConfigurationItemsToApplicationCmd)
}
