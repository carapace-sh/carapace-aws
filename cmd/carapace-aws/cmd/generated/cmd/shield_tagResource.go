package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a resource in Shield.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_tagResourceCmd).Standalone()

		shield_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add or update tags for.")
		shield_tagResourceCmd.Flags().String("tags", "", "The tags that you want to modify or add to the resource.")
		shield_tagResourceCmd.MarkFlagRequired("resource-arn")
		shield_tagResourceCmd.MarkFlagRequired("tags")
	})
	shieldCmd.AddCommand(shield_tagResourceCmd)
}
