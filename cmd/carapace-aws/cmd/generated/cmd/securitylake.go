package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylakeCmd = &cobra.Command{
	Use:   "securitylake",
	Short: "Amazon Security Lake is a fully managed security data lake service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylakeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylakeCmd).Standalone()

	})
	rootCmd.AddCommand(securitylakeCmd)
}
