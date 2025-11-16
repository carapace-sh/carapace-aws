package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listTagsForResourceCmd).Standalone()

		lexv2Models_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to get a list of tags for.")
		lexv2Models_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listTagsForResourceCmd)
}
