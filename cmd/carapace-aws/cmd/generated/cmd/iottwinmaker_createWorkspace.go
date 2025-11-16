package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createWorkspaceCmd = &cobra.Command{
	Use:   "create-workspace",
	Short: "Creates a workplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_createWorkspaceCmd).Standalone()

		iottwinmaker_createWorkspaceCmd.Flags().String("description", "", "The description of the workspace.")
		iottwinmaker_createWorkspaceCmd.Flags().String("role", "", "The ARN of the execution role associated with the workspace.")
		iottwinmaker_createWorkspaceCmd.Flags().String("s3-location", "", "The ARN of the S3 bucket where resources associated with the workspace are stored.")
		iottwinmaker_createWorkspaceCmd.Flags().String("tags", "", "Metadata that you can use to manage the workspace")
		iottwinmaker_createWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_createWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_createWorkspaceCmd)
}
