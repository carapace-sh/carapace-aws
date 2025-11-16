package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateImagePermissionsCmd = &cobra.Command{
	Use:   "update-image-permissions",
	Short: "Adds or updates permissions for the specified private image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateImagePermissionsCmd).Standalone()

	appstream_updateImagePermissionsCmd.Flags().String("image-permissions", "", "The permissions for the image.")
	appstream_updateImagePermissionsCmd.Flags().String("name", "", "The name of the private image.")
	appstream_updateImagePermissionsCmd.Flags().String("shared-account-id", "", "The 12-digit identifier of the AWS account for which you want add or update image permissions.")
	appstream_updateImagePermissionsCmd.MarkFlagRequired("image-permissions")
	appstream_updateImagePermissionsCmd.MarkFlagRequired("name")
	appstream_updateImagePermissionsCmd.MarkFlagRequired("shared-account-id")
	appstreamCmd.AddCommand(appstream_updateImagePermissionsCmd)
}
