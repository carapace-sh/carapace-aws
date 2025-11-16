package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_importCustomWorkspaceImageCmd = &cobra.Command{
	Use:   "import-custom-workspace-image",
	Short: "Imports the specified Windows 10 or 11 Bring Your Own License (BYOL) image into Amazon WorkSpaces using EC2 Image Builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_importCustomWorkspaceImageCmd).Standalone()

	workspaces_importCustomWorkspaceImageCmd.Flags().String("compute-type", "", "The supported compute type for the WorkSpace image.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("image-description", "", "The description of the WorkSpace image.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("image-name", "", "The name of the WorkSpace image.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("image-source", "", "The options for image import source.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("infrastructure-configuration-arn", "", "The infrastructure configuration ARN that specifies how the WorkSpace image is built.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("os-version", "", "The OS version for the WorkSpace image source.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("platform", "", "The platform for the WorkSpace image source.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("protocol", "", "The supported protocol for the WorkSpace image.")
	workspaces_importCustomWorkspaceImageCmd.Flags().String("tags", "", "The resource tags.")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("compute-type")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("image-description")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("image-name")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("image-source")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("infrastructure-configuration-arn")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("os-version")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("platform")
	workspaces_importCustomWorkspaceImageCmd.MarkFlagRequired("protocol")
	workspacesCmd.AddCommand(workspaces_importCustomWorkspaceImageCmd)
}
