package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a given tag to a resource in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_tagResourceCmd).Standalone()

		lookoutequipment_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the specific resource to which the tag should be associated.")
		lookoutequipment_tagResourceCmd.Flags().String("tags", "", "The tag or tags to be associated with a specific resource.")
		lookoutequipment_tagResourceCmd.MarkFlagRequired("resource-arn")
		lookoutequipment_tagResourceCmd.MarkFlagRequired("tags")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_tagResourceCmd)
}
