package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of tags with a report definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_tagResourceCmd).Standalone()

	cur_tagResourceCmd.Flags().String("report-name", "", "The report name of the report definition that tags are to be associated with.")
	cur_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the report definition resource.")
	cur_tagResourceCmd.MarkFlagRequired("report-name")
	cur_tagResourceCmd.MarkFlagRequired("tags")
	curCmd.AddCommand(cur_tagResourceCmd)
}
