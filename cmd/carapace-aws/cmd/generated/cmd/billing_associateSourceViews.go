package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_associateSourceViewsCmd = &cobra.Command{
	Use:   "associate-source-views",
	Short: "Associates one or more source billing views with an existing billing view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_associateSourceViewsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billing_associateSourceViewsCmd).Standalone()

		billing_associateSourceViewsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing view to associate source views with.")
		billing_associateSourceViewsCmd.Flags().String("source-views", "", "A list of ARNs of the source billing views to associate.")
		billing_associateSourceViewsCmd.MarkFlagRequired("arn")
		billing_associateSourceViewsCmd.MarkFlagRequired("source-views")
	})
	billingCmd.AddCommand(billing_associateSourceViewsCmd)
}
