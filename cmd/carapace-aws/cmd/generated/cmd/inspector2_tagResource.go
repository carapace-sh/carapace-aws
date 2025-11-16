package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_tagResourceCmd).Standalone()

		inspector2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to apply a tag to.")
		inspector2_tagResourceCmd.Flags().String("tags", "", "The tags to be added to a resource.")
		inspector2_tagResourceCmd.MarkFlagRequired("resource-arn")
		inspector2_tagResourceCmd.MarkFlagRequired("tags")
	})
	inspector2Cmd.AddCommand(inspector2_tagResourceCmd)
}
