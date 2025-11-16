package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getInstanceCmd = &cobra.Command{
	Use:   "get-instance",
	Short: "Gets information about a specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_getInstanceCmd).Standalone()

		servicediscovery_getInstanceCmd.Flags().String("instance-id", "", "The ID of the instance that you want to get information about.")
		servicediscovery_getInstanceCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that the instance is associated with.")
		servicediscovery_getInstanceCmd.MarkFlagRequired("instance-id")
		servicediscovery_getInstanceCmd.MarkFlagRequired("service-id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_getInstanceCmd)
}
