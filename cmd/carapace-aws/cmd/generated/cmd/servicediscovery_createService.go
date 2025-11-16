package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Creates a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_createServiceCmd).Standalone()

	servicediscovery_createServiceCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed `CreateService` requests to be retried without the risk of running the operation twice.")
	servicediscovery_createServiceCmd.Flags().String("description", "", "A description for the service.")
	servicediscovery_createServiceCmd.Flags().String("dns-config", "", "A complex type that contains information about the Amazon Route\u00a053 records that you want Cloud Map to create when you register an instance.")
	servicediscovery_createServiceCmd.Flags().String("health-check-config", "", "*Public DNS and HTTP namespaces only.* A complex type that contains settings for an optional Route\u00a053 health check.")
	servicediscovery_createServiceCmd.Flags().String("health-check-custom-config", "", "A complex type that contains information about an optional custom health check.")
	servicediscovery_createServiceCmd.Flags().String("name", "", "The name that you want to assign to the service.")
	servicediscovery_createServiceCmd.Flags().String("namespace-id", "", "The ID or Amazon Resource Name (ARN) of the namespace that you want to use to create the service.")
	servicediscovery_createServiceCmd.Flags().String("tags", "", "The tags to add to the service.")
	servicediscovery_createServiceCmd.Flags().String("type", "", "If present, specifies that the service instances are only discoverable using the `DiscoverInstances` API operation.")
	servicediscovery_createServiceCmd.MarkFlagRequired("name")
	servicediscoveryCmd.AddCommand(servicediscovery_createServiceCmd)
}
