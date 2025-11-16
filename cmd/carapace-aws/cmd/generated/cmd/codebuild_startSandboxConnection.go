package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_startSandboxConnectionCmd = &cobra.Command{
	Use:   "start-sandbox-connection",
	Short: "Starts a sandbox connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_startSandboxConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_startSandboxConnectionCmd).Standalone()

		codebuild_startSandboxConnectionCmd.Flags().String("sandbox-id", "", "A `sandboxId` or `sandboxArn`.")
		codebuild_startSandboxConnectionCmd.MarkFlagRequired("sandbox-id")
	})
	codebuildCmd.AddCommand(codebuild_startSandboxConnectionCmd)
}
