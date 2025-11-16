package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_putMeteredProductCmd = &cobra.Command{
	Use:   "put-metered-product",
	Short: "Adds a metered product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_putMeteredProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_putMeteredProductCmd).Standalone()

		deadline_putMeteredProductCmd.Flags().String("license-endpoint-id", "", "The license endpoint ID to add to the metered product.")
		deadline_putMeteredProductCmd.Flags().String("product-id", "", "The product ID to add to the metered product.")
		deadline_putMeteredProductCmd.MarkFlagRequired("license-endpoint-id")
		deadline_putMeteredProductCmd.MarkFlagRequired("product-id")
	})
	deadlineCmd.AddCommand(deadline_putMeteredProductCmd)
}
