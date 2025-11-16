package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_createPublicDnsNamespaceCmd = &cobra.Command{
	Use:   "create-public-dns-namespace",
	Short: "Creates a public namespace based on DNS, which is visible on the internet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_createPublicDnsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_createPublicDnsNamespaceCmd).Standalone()

		servicediscovery_createPublicDnsNamespaceCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed `CreatePublicDnsNamespace` requests to be retried without the risk of running the operation twice.")
		servicediscovery_createPublicDnsNamespaceCmd.Flags().String("description", "", "A description for the namespace.")
		servicediscovery_createPublicDnsNamespaceCmd.Flags().String("name", "", "The name that you want to assign to this namespace.")
		servicediscovery_createPublicDnsNamespaceCmd.Flags().String("properties", "", "Properties for the public DNS namespace.")
		servicediscovery_createPublicDnsNamespaceCmd.Flags().String("tags", "", "The tags to add to the namespace.")
		servicediscovery_createPublicDnsNamespaceCmd.MarkFlagRequired("name")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_createPublicDnsNamespaceCmd)
}
