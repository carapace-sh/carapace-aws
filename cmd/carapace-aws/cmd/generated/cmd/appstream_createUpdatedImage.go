package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createUpdatedImageCmd = &cobra.Command{
	Use:   "create-updated-image",
	Short: "Creates a new image with the latest Windows operating system updates, driver updates, and AppStream 2.0 agent software.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createUpdatedImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createUpdatedImageCmd).Standalone()

		appstream_createUpdatedImageCmd.Flags().Bool("dry-run", false, "Indicates whether to display the status of image update availability before AppStream 2.0 initiates the process of creating a new updated image.")
		appstream_createUpdatedImageCmd.Flags().String("existing-image-name", "", "The name of the image to update.")
		appstream_createUpdatedImageCmd.Flags().String("new-image-description", "", "The description to display for the new image.")
		appstream_createUpdatedImageCmd.Flags().String("new-image-display-name", "", "The name to display for the new image.")
		appstream_createUpdatedImageCmd.Flags().String("new-image-name", "", "The name of the new image.")
		appstream_createUpdatedImageCmd.Flags().String("new-image-tags", "", "The tags to associate with the new image.")
		appstream_createUpdatedImageCmd.Flags().Bool("no-dry-run", false, "Indicates whether to display the status of image update availability before AppStream 2.0 initiates the process of creating a new updated image.")
		appstream_createUpdatedImageCmd.MarkFlagRequired("existing-image-name")
		appstream_createUpdatedImageCmd.MarkFlagRequired("new-image-name")
		appstream_createUpdatedImageCmd.Flag("no-dry-run").Hidden = true
	})
	appstreamCmd.AddCommand(appstream_createUpdatedImageCmd)
}
