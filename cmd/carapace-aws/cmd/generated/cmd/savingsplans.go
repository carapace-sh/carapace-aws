package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplansCmd = &cobra.Command{
	Use:   "savingsplans",
	Short: "Savings Plans are a pricing model that offer significant savings on Amazon Web Services usage (for example, on Amazon EC2 instances).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplansCmd).Standalone()

	})
	rootCmd.AddCommand(savingsplansCmd)
}
