package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags that are associated with for an App Runner resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listTagsForResourceCmd).Standalone()

	apprunner_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that a tag list is requested for.")
	apprunner_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	apprunnerCmd.AddCommand(apprunner_listTagsForResourceCmd)
}
