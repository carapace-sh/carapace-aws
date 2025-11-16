package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_associateAppBlockBuilderAppBlockCmd = &cobra.Command{
	Use:   "associate-app-block-builder-app-block",
	Short: "Associates the specified app block builder with the specified app block.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_associateAppBlockBuilderAppBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_associateAppBlockBuilderAppBlockCmd).Standalone()

		appstream_associateAppBlockBuilderAppBlockCmd.Flags().String("app-block-arn", "", "The ARN of the app block.")
		appstream_associateAppBlockBuilderAppBlockCmd.Flags().String("app-block-builder-name", "", "The name of the app block builder.")
		appstream_associateAppBlockBuilderAppBlockCmd.MarkFlagRequired("app-block-arn")
		appstream_associateAppBlockBuilderAppBlockCmd.MarkFlagRequired("app-block-builder-name")
	})
	appstreamCmd.AddCommand(appstream_associateAppBlockBuilderAppBlockCmd)
}
