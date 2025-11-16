package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeAppBlocksCmd = &cobra.Command{
	Use:   "describe-app-blocks",
	Short: "Retrieves a list that describes one or more app blocks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeAppBlocksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeAppBlocksCmd).Standalone()

		appstream_describeAppBlocksCmd.Flags().String("arns", "", "The ARNs of the app blocks.")
		appstream_describeAppBlocksCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeAppBlocksCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeAppBlocksCmd)
}
