package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeEksAnywhereSubscriptionCmd = &cobra.Command{
	Use:   "describe-eks-anywhere-subscription",
	Short: "Returns descriptive information about a subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeEksAnywhereSubscriptionCmd).Standalone()

	eks_describeEksAnywhereSubscriptionCmd.Flags().String("id", "", "The ID of the subscription.")
	eks_describeEksAnywhereSubscriptionCmd.MarkFlagRequired("id")
	eksCmd.AddCommand(eks_describeEksAnywhereSubscriptionCmd)
}
