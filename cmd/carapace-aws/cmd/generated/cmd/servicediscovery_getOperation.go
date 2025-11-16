package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getOperationCmd = &cobra.Command{
	Use:   "get-operation",
	Short: "Gets information about any operation that returns an operation ID in the response, such as a `CreateHttpNamespace` request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_getOperationCmd).Standalone()

		servicediscovery_getOperationCmd.Flags().String("operation-id", "", "The ID of the operation that you want to get more information about.")
		servicediscovery_getOperationCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the namespace associated with the operation, as specified in the namespace `ResourceOwner` field.")
		servicediscovery_getOperationCmd.MarkFlagRequired("operation-id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_getOperationCmd)
}
