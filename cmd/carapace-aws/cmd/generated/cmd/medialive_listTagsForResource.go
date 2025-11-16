package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Produces list of tags that have been created for a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listTagsForResourceCmd).Standalone()

		medialive_listTagsForResourceCmd.Flags().String("resource-arn", "", "")
		medialive_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	medialiveCmd.AddCommand(medialive_listTagsForResourceCmd)
}
