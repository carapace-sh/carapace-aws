package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeImageBuildersCmd = &cobra.Command{
	Use:   "describe-image-builders",
	Short: "Retrieves a list that describes one or more specified image builders, if the image builder names are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeImageBuildersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeImageBuildersCmd).Standalone()

		appstream_describeImageBuildersCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeImageBuildersCmd.Flags().String("names", "", "The names of the image builders to describe.")
		appstream_describeImageBuildersCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeImageBuildersCmd)
}
