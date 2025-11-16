package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2RuntimeCmd = &cobra.Command{
	Use:   "lexv2-runtime",
	Short: "This section contains documentation for the Amazon Lex V2 Runtime V2 API operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2RuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2RuntimeCmd).Standalone()

	})
	rootCmd.AddCommand(lexv2RuntimeCmd)
}
