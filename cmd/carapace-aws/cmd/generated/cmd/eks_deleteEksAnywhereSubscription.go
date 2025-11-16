package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteEksAnywhereSubscriptionCmd = &cobra.Command{
	Use:   "delete-eks-anywhere-subscription",
	Short: "Deletes an expired or inactive subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteEksAnywhereSubscriptionCmd).Standalone()

	eks_deleteEksAnywhereSubscriptionCmd.Flags().String("id", "", "The ID of the subscription.")
	eks_deleteEksAnywhereSubscriptionCmd.MarkFlagRequired("id")
	eksCmd.AddCommand(eks_deleteEksAnywhereSubscriptionCmd)
}
