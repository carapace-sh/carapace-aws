package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Retrieves a list of configuration items that have tags as specified by the key-value pairs, name and value, passed to the optional parameter `filters`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_describeTagsCmd).Standalone()

		discovery_describeTagsCmd.Flags().String("filters", "", "You can filter the list using a *key*-*value* format.")
		discovery_describeTagsCmd.Flags().String("max-results", "", "The total number of items to return in a single page of output.")
		discovery_describeTagsCmd.Flags().String("next-token", "", "A token to start the list.")
	})
	discoveryCmd.AddCommand(discovery_describeTagsCmd)
}
