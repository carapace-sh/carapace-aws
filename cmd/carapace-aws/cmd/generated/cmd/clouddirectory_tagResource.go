package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "An API operation for adding tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_tagResourceCmd).Standalone()

	clouddirectory_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	clouddirectory_tagResourceCmd.Flags().String("tags", "", "A list of tag key-value pairs.")
	clouddirectory_tagResourceCmd.MarkFlagRequired("resource-arn")
	clouddirectory_tagResourceCmd.MarkFlagRequired("tags")
	clouddirectoryCmd.AddCommand(clouddirectory_tagResourceCmd)
}
