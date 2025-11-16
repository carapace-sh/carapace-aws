package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_invalidateProjectCacheCmd = &cobra.Command{
	Use:   "invalidate-project-cache",
	Short: "Resets the cache for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_invalidateProjectCacheCmd).Standalone()

	codebuild_invalidateProjectCacheCmd.Flags().String("project-name", "", "The name of the CodeBuild build project that the cache is reset for.")
	codebuild_invalidateProjectCacheCmd.MarkFlagRequired("project-name")
	codebuildCmd.AddCommand(codebuild_invalidateProjectCacheCmd)
}
