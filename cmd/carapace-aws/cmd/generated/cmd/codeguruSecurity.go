package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurityCmd = &cobra.Command{
	Use:   "codeguru-security",
	Short: "On November 20, 2025, AWS will discontinue support for Amazon CodeGuru Security.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurityCmd).Standalone()

	})
	rootCmd.AddCommand(codeguruSecurityCmd)
}
