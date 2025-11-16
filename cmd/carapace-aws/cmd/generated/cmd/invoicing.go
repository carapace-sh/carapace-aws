package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicingCmd = &cobra.Command{
	Use:   "invoicing",
	Short: "**Amazon Web Services Invoice Configuration**\n\nYou can use Amazon Web Services Invoice Configuration APIs to programmatically create, update, delete, get, and list invoice units.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicingCmd).Standalone()

	})
	rootCmd.AddCommand(invoicingCmd)
}
