package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updatePrivateDnsNamespaceCmd = &cobra.Command{
	Use:   "update-private-dns-namespace",
	Short: "Updates a private DNS namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updatePrivateDnsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updatePrivateDnsNamespaceCmd).Standalone()

		servicediscovery_updatePrivateDnsNamespaceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the namespace that you want to update.")
		servicediscovery_updatePrivateDnsNamespaceCmd.Flags().String("namespace", "", "Updated properties for the private DNS namespace.")
		servicediscovery_updatePrivateDnsNamespaceCmd.Flags().String("updater-request-id", "", "A unique string that identifies the request and that allows failed `UpdatePrivateDnsNamespace` requests to be retried without the risk of running the operation twice.")
		servicediscovery_updatePrivateDnsNamespaceCmd.MarkFlagRequired("id")
		servicediscovery_updatePrivateDnsNamespaceCmd.MarkFlagRequired("namespace")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updatePrivateDnsNamespaceCmd)
}
