package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeAppBlockBuilderAppBlockAssociationsCmd = &cobra.Command{
	Use:   "describe-app-block-builder-app-block-associations",
	Short: "Retrieves a list that describes one or more app block builder associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeAppBlockBuilderAppBlockAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeAppBlockBuilderAppBlockAssociationsCmd).Standalone()

		appstream_describeAppBlockBuilderAppBlockAssociationsCmd.Flags().String("app-block-arn", "", "The ARN of the app block.")
		appstream_describeAppBlockBuilderAppBlockAssociationsCmd.Flags().String("app-block-builder-name", "", "The name of the app block builder.")
		appstream_describeAppBlockBuilderAppBlockAssociationsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeAppBlockBuilderAppBlockAssociationsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeAppBlockBuilderAppBlockAssociationsCmd)
}
