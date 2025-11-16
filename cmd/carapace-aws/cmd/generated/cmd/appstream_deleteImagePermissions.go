package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteImagePermissionsCmd = &cobra.Command{
	Use:   "delete-image-permissions",
	Short: "Deletes permissions for the specified private image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteImagePermissionsCmd).Standalone()

	appstream_deleteImagePermissionsCmd.Flags().String("name", "", "The name of the private image.")
	appstream_deleteImagePermissionsCmd.Flags().String("shared-account-id", "", "The 12-digit identifier of the AWS account for which to delete image permissions.")
	appstream_deleteImagePermissionsCmd.MarkFlagRequired("name")
	appstream_deleteImagePermissionsCmd.MarkFlagRequired("shared-account-id")
	appstreamCmd.AddCommand(appstream_deleteImagePermissionsCmd)
}
