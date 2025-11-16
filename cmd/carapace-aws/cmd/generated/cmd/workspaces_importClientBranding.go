package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_importClientBrandingCmd = &cobra.Command{
	Use:   "import-client-branding",
	Short: "Imports client branding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_importClientBrandingCmd).Standalone()

	workspaces_importClientBrandingCmd.Flags().String("device-type-android", "", "The branding information to import for Android devices.")
	workspaces_importClientBrandingCmd.Flags().String("device-type-ios", "", "The branding information to import for iOS devices.")
	workspaces_importClientBrandingCmd.Flags().String("device-type-linux", "", "The branding information to import for Linux devices.")
	workspaces_importClientBrandingCmd.Flags().String("device-type-osx", "", "The branding information to import for macOS devices.")
	workspaces_importClientBrandingCmd.Flags().String("device-type-web", "", "The branding information to import for web access.")
	workspaces_importClientBrandingCmd.Flags().String("device-type-windows", "", "The branding information to import for Windows devices.")
	workspaces_importClientBrandingCmd.Flags().String("resource-id", "", "The directory identifier of the WorkSpace for which you want to import client branding.")
	workspaces_importClientBrandingCmd.MarkFlagRequired("resource-id")
	workspacesCmd.AddCommand(workspaces_importClientBrandingCmd)
}
