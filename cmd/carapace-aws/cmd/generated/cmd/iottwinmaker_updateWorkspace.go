package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_updateWorkspaceCmd = &cobra.Command{
	Use:   "update-workspace",
	Short: "Updates a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_updateWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_updateWorkspaceCmd).Standalone()

		iottwinmaker_updateWorkspaceCmd.Flags().String("description", "", "The description of the workspace.")
		iottwinmaker_updateWorkspaceCmd.Flags().String("role", "", "The ARN of the execution role associated with the workspace.")
		iottwinmaker_updateWorkspaceCmd.Flags().String("s3-location", "", "The ARN of the S3 bucket where resources associated with the workspace are stored.")
		iottwinmaker_updateWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_updateWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_updateWorkspaceCmd)
}
