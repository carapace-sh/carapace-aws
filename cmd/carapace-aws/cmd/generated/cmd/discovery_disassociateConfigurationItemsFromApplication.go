package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_disassociateConfigurationItemsFromApplicationCmd = &cobra.Command{
	Use:   "disassociate-configuration-items-from-application",
	Short: "Disassociates one or more configuration items from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_disassociateConfigurationItemsFromApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_disassociateConfigurationItemsFromApplicationCmd).Standalone()

		discovery_disassociateConfigurationItemsFromApplicationCmd.Flags().String("application-configuration-id", "", "Configuration ID of an application from which each item is disassociated.")
		discovery_disassociateConfigurationItemsFromApplicationCmd.Flags().String("configuration-ids", "", "Configuration ID of each item to be disassociated from an application.")
		discovery_disassociateConfigurationItemsFromApplicationCmd.MarkFlagRequired("application-configuration-id")
		discovery_disassociateConfigurationItemsFromApplicationCmd.MarkFlagRequired("configuration-ids")
	})
	discoveryCmd.AddCommand(discovery_disassociateConfigurationItemsFromApplicationCmd)
}
