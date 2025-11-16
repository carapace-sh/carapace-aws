package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerCmd = &cobra.Command{
	Use:   "sagemaker",
	Short: "Provides APIs for creating and managing SageMaker resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerCmd).Standalone()

	rootCmd.AddCommand(sagemakerCmd)
}
