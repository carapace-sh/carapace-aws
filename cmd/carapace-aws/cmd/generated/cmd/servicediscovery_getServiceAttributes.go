package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getServiceAttributesCmd = &cobra.Command{
	Use:   "get-service-attributes",
	Short: "Returns the attributes associated with a specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getServiceAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_getServiceAttributesCmd).Standalone()

		servicediscovery_getServiceAttributesCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to get attributes for.")
		servicediscovery_getServiceAttributesCmd.MarkFlagRequired("service-id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_getServiceAttributesCmd)
}
