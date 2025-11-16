package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_deregisterInstanceCmd = &cobra.Command{
	Use:   "deregister-instance",
	Short: "Deletes the Amazon Route\u00a053 DNS records and health check, if any, that Cloud Map created for the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_deregisterInstanceCmd).Standalone()

	servicediscovery_deregisterInstanceCmd.Flags().String("instance-id", "", "The value that you specified for `Id` in the [RegisterInstance](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html) request.")
	servicediscovery_deregisterInstanceCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that the instance is associated with.")
	servicediscovery_deregisterInstanceCmd.MarkFlagRequired("instance-id")
	servicediscovery_deregisterInstanceCmd.MarkFlagRequired("service-id")
	servicediscoveryCmd.AddCommand(servicediscovery_deregisterInstanceCmd)
}
