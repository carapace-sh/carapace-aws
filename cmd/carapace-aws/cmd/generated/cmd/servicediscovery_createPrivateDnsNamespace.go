package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_createPrivateDnsNamespaceCmd = &cobra.Command{
	Use:   "create-private-dns-namespace",
	Short: "Creates a private namespace based on DNS, which is visible only inside a specified Amazon VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_createPrivateDnsNamespaceCmd).Standalone()

	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed `CreatePrivateDnsNamespace` requests to be retried without the risk of running the operation twice.")
	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("description", "", "A description for the namespace.")
	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("name", "", "The name that you want to assign to this namespace.")
	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("properties", "", "Properties for the private DNS namespace.")
	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("tags", "", "The tags to add to the namespace.")
	servicediscovery_createPrivateDnsNamespaceCmd.Flags().String("vpc", "", "The ID of the Amazon VPC that you want to associate the namespace with.")
	servicediscovery_createPrivateDnsNamespaceCmd.MarkFlagRequired("name")
	servicediscovery_createPrivateDnsNamespaceCmd.MarkFlagRequired("vpc")
	servicediscoveryCmd.AddCommand(servicediscovery_createPrivateDnsNamespaceCmd)
}
