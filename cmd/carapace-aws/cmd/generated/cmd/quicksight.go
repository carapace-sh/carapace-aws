package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksightCmd = &cobra.Command{
	Use:   "quicksight",
	Short: "Amazon Quick Suite API Reference\n\nAmazon Quick Sight is a fully managed, serverless business intelligence service for the Amazon Web Services Cloud that makes it easy to extend data and insights to every user in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksightCmd).Standalone()

	rootCmd.AddCommand(quicksightCmd)
}
