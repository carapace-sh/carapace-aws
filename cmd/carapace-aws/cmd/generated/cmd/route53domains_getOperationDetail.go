package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_getOperationDetailCmd = &cobra.Command{
	Use:   "get-operation-detail",
	Short: "This operation returns the current status of an operation that is not completed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_getOperationDetailCmd).Standalone()

	route53domains_getOperationDetailCmd.Flags().String("operation-id", "", "The identifier for the operation for which you want to get the status.")
	route53domains_getOperationDetailCmd.MarkFlagRequired("operation-id")
	route53domainsCmd.AddCommand(route53domains_getOperationDetailCmd)
}
