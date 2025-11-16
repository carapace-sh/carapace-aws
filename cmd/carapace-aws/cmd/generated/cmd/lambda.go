package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "Lambda",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambdaCmd).Standalone()

	})
	rootCmd.AddCommand(lambdaCmd)
}
