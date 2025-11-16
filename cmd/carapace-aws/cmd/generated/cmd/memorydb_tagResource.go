package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Use this operation to add tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_tagResourceCmd).Standalone()

		memorydb_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which the tags are to be added.")
		memorydb_tagResourceCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		memorydb_tagResourceCmd.MarkFlagRequired("resource-arn")
		memorydb_tagResourceCmd.MarkFlagRequired("tags")
	})
	memorydbCmd.AddCommand(memorydb_tagResourceCmd)
}
