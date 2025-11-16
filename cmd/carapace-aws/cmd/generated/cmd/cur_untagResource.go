package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cur_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates a set of tags from a report definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cur_untagResourceCmd).Standalone()

	cur_untagResourceCmd.Flags().String("report-name", "", "The report name of the report definition that tags are to be disassociated from.")
	cur_untagResourceCmd.Flags().String("tag-keys", "", "The tags to be disassociated from the report definition resource.")
	cur_untagResourceCmd.MarkFlagRequired("report-name")
	cur_untagResourceCmd.MarkFlagRequired("tag-keys")
	curCmd.AddCommand(cur_untagResourceCmd)
}
