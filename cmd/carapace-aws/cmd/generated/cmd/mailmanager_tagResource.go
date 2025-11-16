package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags (keys and values) to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_tagResourceCmd).Standalone()

		mailmanager_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
		mailmanager_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
		mailmanager_tagResourceCmd.MarkFlagRequired("resource-arn")
		mailmanager_tagResourceCmd.MarkFlagRequired("tags")
	})
	mailmanagerCmd.AddCommand(mailmanager_tagResourceCmd)
}
