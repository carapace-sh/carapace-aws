package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updateServiceAttributesCmd = &cobra.Command{
	Use:   "update-service-attributes",
	Short: "Submits a request to update a specified service to add service-level attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updateServiceAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updateServiceAttributesCmd).Standalone()

		servicediscovery_updateServiceAttributesCmd.Flags().String("attributes", "", "A string map that contains attribute key-value pairs.")
		servicediscovery_updateServiceAttributesCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to update.")
		servicediscovery_updateServiceAttributesCmd.MarkFlagRequired("attributes")
		servicediscovery_updateServiceAttributesCmd.MarkFlagRequired("service-id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updateServiceAttributesCmd)
}
