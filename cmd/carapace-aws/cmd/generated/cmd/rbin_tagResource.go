package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns tags to the specified retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rbin_tagResourceCmd).Standalone()

		rbin_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the retention rule.")
		rbin_tagResourceCmd.Flags().String("tags", "", "Information about the tags to assign to the retention rule.")
		rbin_tagResourceCmd.MarkFlagRequired("resource-arn")
		rbin_tagResourceCmd.MarkFlagRequired("tags")
	})
	rbinCmd.AddCommand(rbin_tagResourceCmd)
}
