package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a specific tag from a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_untagResourceCmd).Standalone()

		lookoutequipment_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which the tag is currently associated.")
		lookoutequipment_untagResourceCmd.Flags().String("tag-keys", "", "Specifies the key of the tag to be removed from a specified resource.")
		lookoutequipment_untagResourceCmd.MarkFlagRequired("resource-arn")
		lookoutequipment_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_untagResourceCmd)
}
