package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_discoverInstancesCmd = &cobra.Command{
	Use:   "discover-instances",
	Short: "Discovers registered instances for a specified namespace and service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_discoverInstancesCmd).Standalone()

	servicediscovery_discoverInstancesCmd.Flags().String("health-status", "", "The health status of the instances that you want to discover.")
	servicediscovery_discoverInstancesCmd.Flags().String("max-results", "", "The maximum number of instances that you want Cloud Map to return in the response to a `DiscoverInstances` request.")
	servicediscovery_discoverInstancesCmd.Flags().String("namespace-name", "", "The `HttpName` name of the namespace.")
	servicediscovery_discoverInstancesCmd.Flags().String("optional-parameters", "", "Opportunistic filters to scope the results based on custom attributes.")
	servicediscovery_discoverInstancesCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the namespace associated with the instance, as specified in the namespace `ResourceOwner` field.")
	servicediscovery_discoverInstancesCmd.Flags().String("query-parameters", "", "Filters to scope the results based on custom attributes for the instance (for example, `{version=v1, az=1a}`).")
	servicediscovery_discoverInstancesCmd.Flags().String("service-name", "", "The name of the service that you specified when you registered the instance.")
	servicediscovery_discoverInstancesCmd.MarkFlagRequired("namespace-name")
	servicediscovery_discoverInstancesCmd.MarkFlagRequired("service-name")
	servicediscoveryCmd.AddCommand(servicediscovery_discoverInstancesCmd)
}
