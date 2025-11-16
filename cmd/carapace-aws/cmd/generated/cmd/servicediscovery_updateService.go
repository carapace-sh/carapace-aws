package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updateServiceCmd = &cobra.Command{
	Use:   "update-service",
	Short: "Submits a request to perform the following operations:\n\n- Update the TTL setting for existing `DnsRecords` configurations\n- Add, update, or delete `HealthCheckConfig` for a specified service\n  \n  You can't add, update, or delete a `HealthCheckCustomConfig` configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updateServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updateServiceCmd).Standalone()

		servicediscovery_updateServiceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to update.")
		servicediscovery_updateServiceCmd.Flags().String("service", "", "A complex type that contains the new settings for the service.")
		servicediscovery_updateServiceCmd.MarkFlagRequired("id")
		servicediscovery_updateServiceCmd.MarkFlagRequired("service")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updateServiceCmd)
}
