package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_tagResourceCmd).Standalone()

	outposts_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	outposts_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	outposts_tagResourceCmd.MarkFlagRequired("resource-arn")
	outposts_tagResourceCmd.MarkFlagRequired("tags")
	outpostsCmd.AddCommand(outposts_tagResourceCmd)
}
