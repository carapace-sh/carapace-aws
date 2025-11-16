package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteMeteredProductCmd = &cobra.Command{
	Use:   "delete-metered-product",
	Short: "Deletes a metered product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteMeteredProductCmd).Standalone()

	deadline_deleteMeteredProductCmd.Flags().String("license-endpoint-id", "", "The ID of the license endpoint from which to remove the metered product.")
	deadline_deleteMeteredProductCmd.Flags().String("product-id", "", "The product ID to remove from the license endpoint.")
	deadline_deleteMeteredProductCmd.MarkFlagRequired("license-endpoint-id")
	deadline_deleteMeteredProductCmd.MarkFlagRequired("product-id")
	deadlineCmd.AddCommand(deadline_deleteMeteredProductCmd)
}
