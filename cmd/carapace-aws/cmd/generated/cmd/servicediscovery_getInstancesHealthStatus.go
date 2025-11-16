package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getInstancesHealthStatusCmd = &cobra.Command{
	Use:   "get-instances-health-status",
	Short: "Gets the current health status (`Healthy`, `Unhealthy`, or `Unknown`) of one or more instances that are associated with a specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getInstancesHealthStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_getInstancesHealthStatusCmd).Standalone()

		servicediscovery_getInstancesHealthStatusCmd.Flags().String("instances", "", "An array that contains the IDs of all the instances that you want to get the health status for.")
		servicediscovery_getInstancesHealthStatusCmd.Flags().String("max-results", "", "The maximum number of instances that you want Cloud Map to return in the response to a `GetInstancesHealthStatus` request.")
		servicediscovery_getInstancesHealthStatusCmd.Flags().String("next-token", "", "For the first `GetInstancesHealthStatus` request, omit this value.")
		servicediscovery_getInstancesHealthStatusCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that the instance is associated with.")
		servicediscovery_getInstancesHealthStatusCmd.MarkFlagRequired("service-id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_getInstancesHealthStatusCmd)
}
