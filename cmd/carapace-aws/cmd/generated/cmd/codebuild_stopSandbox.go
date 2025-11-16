package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_stopSandboxCmd = &cobra.Command{
	Use:   "stop-sandbox",
	Short: "Stops a sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_stopSandboxCmd).Standalone()

	codebuild_stopSandboxCmd.Flags().String("id", "", "Information about the requested sandbox ID.")
	codebuild_stopSandboxCmd.MarkFlagRequired("id")
	codebuildCmd.AddCommand(codebuild_stopSandboxCmd)
}
