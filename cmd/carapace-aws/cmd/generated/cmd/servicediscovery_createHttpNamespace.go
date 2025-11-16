package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_createHttpNamespaceCmd = &cobra.Command{
	Use:   "create-http-namespace",
	Short: "Creates an HTTP namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_createHttpNamespaceCmd).Standalone()

	servicediscovery_createHttpNamespaceCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed `CreateHttpNamespace` requests to be retried without the risk of running the operation twice.")
	servicediscovery_createHttpNamespaceCmd.Flags().String("description", "", "A description for the namespace.")
	servicediscovery_createHttpNamespaceCmd.Flags().String("name", "", "The name that you want to assign to this namespace.")
	servicediscovery_createHttpNamespaceCmd.Flags().String("tags", "", "The tags to add to the namespace.")
	servicediscovery_createHttpNamespaceCmd.MarkFlagRequired("name")
	servicediscoveryCmd.AddCommand(servicediscovery_createHttpNamespaceCmd)
}
