package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_disassociateSourceViewsCmd = &cobra.Command{
	Use:   "disassociate-source-views",
	Short: "Removes the association between one or more source billing views and an existing billing view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_disassociateSourceViewsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billing_disassociateSourceViewsCmd).Standalone()

		billing_disassociateSourceViewsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing view to disassociate source views from.")
		billing_disassociateSourceViewsCmd.Flags().String("source-views", "", "A list of ARNs of the source billing views to disassociate.")
		billing_disassociateSourceViewsCmd.MarkFlagRequired("arn")
		billing_disassociateSourceViewsCmd.MarkFlagRequired("source-views")
	})
	billingCmd.AddCommand(billing_disassociateSourceViewsCmd)
}
