package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricingCmd = &cobra.Command{
	Use:   "pricing",
	Short: "The Amazon Web Services Price List API is a centralized and convenient way to programmatically query Amazon Web Services for services, products, and pricing information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pricingCmd).Standalone()

	})
	rootCmd.AddCommand(pricingCmd)
}
