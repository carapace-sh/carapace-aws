package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_tagResourceCmd).Standalone()

		cleanrooms_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource you want to tag.")
		cleanrooms_tagResourceCmd.Flags().String("tags", "", "A map of objects specifying each key name and value.")
		cleanrooms_tagResourceCmd.MarkFlagRequired("resource-arn")
		cleanrooms_tagResourceCmd.MarkFlagRequired("tags")
	})
	cleanroomsCmd.AddCommand(cleanrooms_tagResourceCmd)
}
