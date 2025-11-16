package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns tags to resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_tagResourceCmd).Standalone()

	emrServerless_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags for.")
	emrServerless_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	emrServerless_tagResourceCmd.MarkFlagRequired("resource-arn")
	emrServerless_tagResourceCmd.MarkFlagRequired("tags")
	emrServerlessCmd.AddCommand(emrServerless_tagResourceCmd)
}
