package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites only the specified tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_tagResourceCmd).Standalone()

		socialmessaging_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
		socialmessaging_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
		socialmessaging_tagResourceCmd.MarkFlagRequired("resource-arn")
		socialmessaging_tagResourceCmd.MarkFlagRequired("tags")
	})
	socialmessagingCmd.AddCommand(socialmessaging_tagResourceCmd)
}
