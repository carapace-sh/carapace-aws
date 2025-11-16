package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_tagResourceCmd).Standalone()

		controltower_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be tagged.")
		controltower_tagResourceCmd.Flags().String("tags", "", "Tags to be applied to the resource.")
		controltower_tagResourceCmd.MarkFlagRequired("resource-arn")
		controltower_tagResourceCmd.MarkFlagRequired("tags")
	})
	controltowerCmd.AddCommand(controltower_tagResourceCmd)
}
