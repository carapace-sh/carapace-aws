package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getDestinationCmd = &cobra.Command{
	Use:   "get-destination",
	Short: "Gets a destination by name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getDestinationCmd).Standalone()

	iotManagedIntegrations_getDestinationCmd.Flags().String("name", "", "The name of the customer-managed destination.")
	iotManagedIntegrations_getDestinationCmd.MarkFlagRequired("name")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getDestinationCmd)
}
