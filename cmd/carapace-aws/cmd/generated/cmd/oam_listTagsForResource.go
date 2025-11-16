package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_listTagsForResourceCmd).Standalone()

		oam_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you want to view tags for.")
		oam_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	oamCmd.AddCommand(oam_listTagsForResourceCmd)
}
