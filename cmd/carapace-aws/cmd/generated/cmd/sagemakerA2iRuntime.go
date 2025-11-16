package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntimeCmd = &cobra.Command{
	Use:   "sagemaker-a2i-runtime",
	Short: "Amazon Augmented AI (Amazon A2I) adds the benefit of human judgment to any machine learning application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntimeCmd).Standalone()

	rootCmd.AddCommand(sagemakerA2iRuntimeCmd)
}
