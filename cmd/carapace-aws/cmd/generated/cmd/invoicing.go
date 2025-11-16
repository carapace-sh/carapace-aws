package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicingCmd = &cobra.Command{
	Use:   "invoicing",
	Short: "**Amazon Web Services Invoice Configuration**",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicingCmd).Standalone()

	})
	rootCmd.AddCommand(invoicingCmd)
}
