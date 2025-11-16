package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_updateApplicationCmd).Standalone()

		appstream_updateApplicationCmd.Flags().String("app-block-arn", "", "The ARN of the app block.")
		appstream_updateApplicationCmd.Flags().String("attributes-to-delete", "", "The attributes to delete for an application.")
		appstream_updateApplicationCmd.Flags().String("description", "", "The description of the application.")
		appstream_updateApplicationCmd.Flags().String("display-name", "", "The display name of the application.")
		appstream_updateApplicationCmd.Flags().String("icon-s3-location", "", "The icon S3 location of the application.")
		appstream_updateApplicationCmd.Flags().String("launch-parameters", "", "The launch parameters of the application.")
		appstream_updateApplicationCmd.Flags().String("launch-path", "", "The launch path of the application.")
		appstream_updateApplicationCmd.Flags().String("name", "", "The name of the application.")
		appstream_updateApplicationCmd.Flags().String("working-directory", "", "The working directory of the application.")
		appstream_updateApplicationCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_updateApplicationCmd)
}
