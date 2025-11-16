package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createApplicationCmd).Standalone()

	appstream_createApplicationCmd.Flags().String("app-block-arn", "", "The app block ARN to which the application should be associated")
	appstream_createApplicationCmd.Flags().String("description", "", "The description of the application.")
	appstream_createApplicationCmd.Flags().String("display-name", "", "The display name of the application.")
	appstream_createApplicationCmd.Flags().String("icon-s3-location", "", "The location in S3 of the application icon.")
	appstream_createApplicationCmd.Flags().String("instance-families", "", "The instance families the application supports.")
	appstream_createApplicationCmd.Flags().String("launch-parameters", "", "The launch parameters of the application.")
	appstream_createApplicationCmd.Flags().String("launch-path", "", "The launch path of the application.")
	appstream_createApplicationCmd.Flags().String("name", "", "The name of the application.")
	appstream_createApplicationCmd.Flags().String("platforms", "", "The platforms the application supports.")
	appstream_createApplicationCmd.Flags().String("tags", "", "The tags assigned to the application.")
	appstream_createApplicationCmd.Flags().String("working-directory", "", "The working directory of the application.")
	appstream_createApplicationCmd.MarkFlagRequired("app-block-arn")
	appstream_createApplicationCmd.MarkFlagRequired("icon-s3-location")
	appstream_createApplicationCmd.MarkFlagRequired("instance-families")
	appstream_createApplicationCmd.MarkFlagRequired("launch-path")
	appstream_createApplicationCmd.MarkFlagRequired("name")
	appstream_createApplicationCmd.MarkFlagRequired("platforms")
	appstreamCmd.AddCommand(appstream_createApplicationCmd)
}
