package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerEdgeCmd = &cobra.Command{
	Use:   "sagemaker-edge",
	Short: "SageMaker Edge Manager dataplane service for communicating with active agents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerEdgeCmd).Standalone()

	rootCmd.AddCommand(sagemakerEdgeCmd)
}
