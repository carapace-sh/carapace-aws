package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Operations for Amazon Web Services Account Management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accountCmd).Standalone()

	})
	rootCmd.AddCommand(accountCmd)
}
