package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_discoverInstancesRevisionCmd = &cobra.Command{
	Use:   "discover-instances-revision",
	Short: "Discovers the increasing revision associated with an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_discoverInstancesRevisionCmd).Standalone()

	servicediscovery_discoverInstancesRevisionCmd.Flags().String("namespace-name", "", "The `HttpName` name of the namespace.")
	servicediscovery_discoverInstancesRevisionCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the namespace associated with the instance, as specified in the namespace `ResourceOwner` field.")
	servicediscovery_discoverInstancesRevisionCmd.Flags().String("service-name", "", "The name of the service that you specified when you registered the instance.")
	servicediscovery_discoverInstancesRevisionCmd.MarkFlagRequired("namespace-name")
	servicediscovery_discoverInstancesRevisionCmd.MarkFlagRequired("service-name")
	servicediscoveryCmd.AddCommand(servicediscovery_discoverInstancesRevisionCmd)
}
