package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceAgreementCmd = &cobra.Command{
	Use:   "marketplace-agreement",
	Short: "AWS Marketplace is a curated digital catalog that customers can use to find, buy, deploy, and manage third-party software, data, and services to build solutions and run their businesses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceAgreementCmd).Standalone()

	rootCmd.AddCommand(marketplaceAgreementCmd)
}
