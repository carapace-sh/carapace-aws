package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns information about the tags applied to this resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listTagsForResourceCmd).Standalone()

		odb_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to list tags for.")
		odb_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	odbCmd.AddCommand(odb_listTagsForResourceCmd)
}
