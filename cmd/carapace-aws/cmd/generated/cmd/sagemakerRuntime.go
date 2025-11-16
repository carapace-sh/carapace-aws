package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerRuntimeCmd = &cobra.Command{
	Use:   "sagemaker-runtime",
	Short: "The Amazon SageMaker AI runtime API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerRuntimeCmd).Standalone()

	})
	rootCmd.AddCommand(sagemakerRuntimeCmd)
}
