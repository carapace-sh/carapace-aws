package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Create tags for a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createTagsCmd).Standalone()

		medialive_createTagsCmd.Flags().String("resource-arn", "", "")
		medialive_createTagsCmd.Flags().String("tags", "", "")
		medialive_createTagsCmd.MarkFlagRequired("resource-arn")
	})
	medialiveCmd.AddCommand(medialive_createTagsCmd)
}
