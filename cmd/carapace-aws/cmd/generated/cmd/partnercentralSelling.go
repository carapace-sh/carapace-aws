package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSellingCmd = &cobra.Command{
	Use:   "partnercentral-selling",
	Short: "AWS Partner Central API for Selling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSellingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSellingCmd).Standalone()

	})
	rootCmd.AddCommand(partnercentralSellingCmd)
}
