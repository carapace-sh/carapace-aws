package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies a tag to the specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_tagResourceCmd).Standalone()

	appflow_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to tag.")
	appflow_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for your flow.")
	appflow_tagResourceCmd.MarkFlagRequired("resource-arn")
	appflow_tagResourceCmd.MarkFlagRequired("tags")
	appflowCmd.AddCommand(appflow_tagResourceCmd)
}
