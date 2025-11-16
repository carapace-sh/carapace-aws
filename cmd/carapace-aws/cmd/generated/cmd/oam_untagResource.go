package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_untagResourceCmd).Standalone()

	oam_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're removing tags from.")
	oam_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	oam_untagResourceCmd.MarkFlagRequired("resource-arn")
	oam_untagResourceCmd.MarkFlagRequired("tag-keys")
	oamCmd.AddCommand(oam_untagResourceCmd)
}
