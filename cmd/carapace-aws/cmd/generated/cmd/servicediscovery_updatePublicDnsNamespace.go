package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updatePublicDnsNamespaceCmd = &cobra.Command{
	Use:   "update-public-dns-namespace",
	Short: "Updates a public DNS namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updatePublicDnsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updatePublicDnsNamespaceCmd).Standalone()

		servicediscovery_updatePublicDnsNamespaceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the namespace being updated.")
		servicediscovery_updatePublicDnsNamespaceCmd.Flags().String("namespace", "", "Updated properties for the public DNS namespace.")
		servicediscovery_updatePublicDnsNamespaceCmd.Flags().String("updater-request-id", "", "A unique string that identifies the request and that allows failed `UpdatePublicDnsNamespace` requests to be retried without the risk of running the operation twice.")
		servicediscovery_updatePublicDnsNamespaceCmd.MarkFlagRequired("id")
		servicediscovery_updatePublicDnsNamespaceCmd.MarkFlagRequired("namespace")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updatePublicDnsNamespaceCmd)
}
