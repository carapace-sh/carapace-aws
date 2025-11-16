package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified Amazon Location Service resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_tagResourceCmd).Standalone()

		location_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags you want to update.")
		location_tagResourceCmd.Flags().String("tags", "", "Applies one or more tags to specific resource.")
		location_tagResourceCmd.MarkFlagRequired("resource-arn")
		location_tagResourceCmd.MarkFlagRequired("tags")
	})
	locationCmd.AddCommand(location_tagResourceCmd)
}
