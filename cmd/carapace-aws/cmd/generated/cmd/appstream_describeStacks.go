package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeStacksCmd = &cobra.Command{
	Use:   "describe-stacks",
	Short: "Retrieves a list that describes one or more specified stacks, if the stack names are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeStacksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeStacksCmd).Standalone()

		appstream_describeStacksCmd.Flags().String("names", "", "The names of the stacks to describe.")
		appstream_describeStacksCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeStacksCmd)
}
