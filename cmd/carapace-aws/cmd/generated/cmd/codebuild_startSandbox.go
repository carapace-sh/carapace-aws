package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_startSandboxCmd = &cobra.Command{
	Use:   "start-sandbox",
	Short: "Starts a sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_startSandboxCmd).Standalone()

	codebuild_startSandboxCmd.Flags().String("idempotency-token", "", "A unique client token.")
	codebuild_startSandboxCmd.Flags().String("project-name", "", "The CodeBuild project name.")
	codebuildCmd.AddCommand(codebuild_startSandboxCmd)
}
