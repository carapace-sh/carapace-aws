package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of tags with a Timestream resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_tagResourceCmd).Standalone()

		timestreamWrite_tagResourceCmd.Flags().String("resource-arn", "", "Identifies the Timestream resource to which tags should be added.")
		timestreamWrite_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Timestream resource.")
		timestreamWrite_tagResourceCmd.MarkFlagRequired("resource-arn")
		timestreamWrite_tagResourceCmd.MarkFlagRequired("tags")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_tagResourceCmd)
}
