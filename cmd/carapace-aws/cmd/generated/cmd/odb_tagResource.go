package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_tagResourceCmd).Standalone()

		odb_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to apply tags to.")
		odb_tagResourceCmd.Flags().String("tags", "", "The list of tags to apply to the resource.")
		odb_tagResourceCmd.MarkFlagRequired("resource-arn")
		odb_tagResourceCmd.MarkFlagRequired("tags")
	})
	odbCmd.AddCommand(odb_tagResourceCmd)
}
