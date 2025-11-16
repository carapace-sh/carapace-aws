package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updateInstanceCustomHealthStatusCmd = &cobra.Command{
	Use:   "update-instance-custom-health-status",
	Short: "Submits a request to change the health status of a custom health check to healthy or unhealthy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updateInstanceCustomHealthStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updateInstanceCustomHealthStatusCmd).Standalone()

		servicediscovery_updateInstanceCustomHealthStatusCmd.Flags().String("instance-id", "", "The ID of the instance that you want to change the health status for.")
		servicediscovery_updateInstanceCustomHealthStatusCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that includes the configuration for the custom health check that you want to change the status for.")
		servicediscovery_updateInstanceCustomHealthStatusCmd.Flags().String("status", "", "The new status of the instance, `HEALTHY` or `UNHEALTHY`.")
		servicediscovery_updateInstanceCustomHealthStatusCmd.MarkFlagRequired("instance-id")
		servicediscovery_updateInstanceCustomHealthStatusCmd.MarkFlagRequired("service-id")
		servicediscovery_updateInstanceCustomHealthStatusCmd.MarkFlagRequired("status")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updateInstanceCustomHealthStatusCmd)
}
