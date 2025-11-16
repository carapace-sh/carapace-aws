package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModelsCmd = &cobra.Command{
	Use:   "lex-models",
	Short: "Amazon Lex Build-Time Actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModelsCmd).Standalone()

	})
	rootCmd.AddCommand(lexModelsCmd)
}
