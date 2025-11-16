package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about Amazon Web Services tags for a specified Amazon Resource Name (ARN) in Shield.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_listTagsForResourceCmd).Standalone()

	shield_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to get tags for.")
	shield_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	shieldCmd.AddCommand(shield_listTagsForResourceCmd)
}
