package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_tagResourceCmd).Standalone()

		oam_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're adding tags to.")
		oam_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		oam_tagResourceCmd.MarkFlagRequired("resource-arn")
		oam_tagResourceCmd.MarkFlagRequired("tags")
	})
	oamCmd.AddCommand(oam_tagResourceCmd)
}
