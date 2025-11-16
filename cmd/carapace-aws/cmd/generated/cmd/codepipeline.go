package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipelineCmd = &cobra.Command{
	Use:   "codepipeline",
	Short: "CodePipeline\n\n**Overview**\n\nThis is the CodePipeline API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipelineCmd).Standalone()

	rootCmd.AddCommand(codepipelineCmd)
}
