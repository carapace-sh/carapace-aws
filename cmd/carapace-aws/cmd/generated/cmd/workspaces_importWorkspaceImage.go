package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_importWorkspaceImageCmd = &cobra.Command{
	Use:   "import-workspace-image",
	Short: "Imports the specified Windows 10 or 11 Bring Your Own License (BYOL) image into Amazon WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_importWorkspaceImageCmd).Standalone()

	workspaces_importWorkspaceImageCmd.Flags().String("applications", "", "If specified, the version of Microsoft Office to subscribe to.")
	workspaces_importWorkspaceImageCmd.Flags().String("ec2-image-id", "", "The identifier of the EC2 image.")
	workspaces_importWorkspaceImageCmd.Flags().String("image-description", "", "The description of the WorkSpace image.")
	workspaces_importWorkspaceImageCmd.Flags().String("image-name", "", "The name of the WorkSpace image.")
	workspaces_importWorkspaceImageCmd.Flags().String("ingestion-process", "", "The ingestion process to be used when importing the image, depending on which protocol you want to use for your BYOL Workspace image, either PCoIP, WSP, or bring your own protocol (BYOP).")
	workspaces_importWorkspaceImageCmd.Flags().String("tags", "", "The tags.")
	workspaces_importWorkspaceImageCmd.MarkFlagRequired("ec2-image-id")
	workspaces_importWorkspaceImageCmd.MarkFlagRequired("image-description")
	workspaces_importWorkspaceImageCmd.MarkFlagRequired("image-name")
	workspaces_importWorkspaceImageCmd.MarkFlagRequired("ingestion-process")
	workspacesCmd.AddCommand(workspaces_importWorkspaceImageCmd)
}
