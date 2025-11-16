package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_updateHttpNamespaceCmd = &cobra.Command{
	Use:   "update-http-namespace",
	Short: "Updates an HTTP namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_updateHttpNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_updateHttpNamespaceCmd).Standalone()

		servicediscovery_updateHttpNamespaceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the namespace that you want to update.")
		servicediscovery_updateHttpNamespaceCmd.Flags().String("namespace", "", "Updated properties for the the HTTP namespace.")
		servicediscovery_updateHttpNamespaceCmd.Flags().String("updater-request-id", "", "A unique string that identifies the request and that allows failed `UpdateHttpNamespace` requests to be retried without the risk of running the operation twice.")
		servicediscovery_updateHttpNamespaceCmd.MarkFlagRequired("id")
		servicediscovery_updateHttpNamespaceCmd.MarkFlagRequired("namespace")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_updateHttpNamespaceCmd)
}
