package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createAppBlockCmd = &cobra.Command{
	Use:   "create-app-block",
	Short: "Creates an app block.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createAppBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createAppBlockCmd).Standalone()

		appstream_createAppBlockCmd.Flags().String("description", "", "The description of the app block.")
		appstream_createAppBlockCmd.Flags().String("display-name", "", "The display name of the app block.")
		appstream_createAppBlockCmd.Flags().String("name", "", "The name of the app block.")
		appstream_createAppBlockCmd.Flags().String("packaging-type", "", "The packaging type of the app block.")
		appstream_createAppBlockCmd.Flags().String("post-setup-script-details", "", "The post setup script details of the app block.")
		appstream_createAppBlockCmd.Flags().String("setup-script-details", "", "The setup script details of the app block.")
		appstream_createAppBlockCmd.Flags().String("source-s3-location", "", "The source S3 location of the app block.")
		appstream_createAppBlockCmd.Flags().String("tags", "", "The tags assigned to the app block.")
		appstream_createAppBlockCmd.MarkFlagRequired("name")
		appstream_createAppBlockCmd.MarkFlagRequired("source-s3-location")
	})
	appstreamCmd.AddCommand(appstream_createAppBlockCmd)
}
