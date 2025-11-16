package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeWorkspaceCmd = &cobra.Command{
	Use:   "describe-workspace",
	Short: "Returns information about an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeWorkspaceCmd).Standalone()

		amp_describeWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to describe.")
		amp_describeWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeWorkspaceCmd)
}
