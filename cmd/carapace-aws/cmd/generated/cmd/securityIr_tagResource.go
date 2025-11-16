package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag(s) to a designated resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_tagResourceCmd).Standalone()

		securityIr_tagResourceCmd.Flags().String("resource-arn", "", "Required element for TagResource to identify the ARN for the resource to add a tag to.")
		securityIr_tagResourceCmd.Flags().String("tags", "", "Required element for ListTagsForResource to provide the content for a tag.")
		securityIr_tagResourceCmd.MarkFlagRequired("resource-arn")
		securityIr_tagResourceCmd.MarkFlagRequired("tags")
	})
	securityIrCmd.AddCommand(securityIr_tagResourceCmd)
}
