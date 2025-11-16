package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associate a set of tags with a Timestream resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_tagResourceCmd).Standalone()

		timestreamQuery_tagResourceCmd.Flags().String("resource-arn", "", "Identifies the Timestream resource to which tags should be added.")
		timestreamQuery_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Timestream resource.")
		timestreamQuery_tagResourceCmd.MarkFlagRequired("resource-arn")
		timestreamQuery_tagResourceCmd.MarkFlagRequired("tags")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_tagResourceCmd)
}
