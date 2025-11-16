package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeAppBlockBuildersCmd = &cobra.Command{
	Use:   "describe-app-block-builders",
	Short: "Retrieves a list that describes one or more app block builders.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeAppBlockBuildersCmd).Standalone()

	appstream_describeAppBlockBuildersCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeAppBlockBuildersCmd.Flags().String("names", "", "The names of the app block builders.")
	appstream_describeAppBlockBuildersCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	appstreamCmd.AddCommand(appstream_describeAppBlockBuildersCmd)
}
