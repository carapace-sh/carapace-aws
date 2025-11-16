package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags associated with the specified report definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_listTagsForResourceCmd).Standalone()

	cur_listTagsForResourceCmd.Flags().String("report-name", "", "The report name of the report definition that tags are to be returned for.")
	cur_listTagsForResourceCmd.MarkFlagRequired("report-name")
	curCmd.AddCommand(cur_listTagsForResourceCmd)
}
