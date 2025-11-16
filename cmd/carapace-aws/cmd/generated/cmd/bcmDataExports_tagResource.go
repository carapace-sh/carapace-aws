package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags for an existing data export definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDataExports_tagResourceCmd).Standalone()

		bcmDataExports_tagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
		bcmDataExports_tagResourceCmd.Flags().String("resource-tags", "", "The tags to associate with the resource.")
		bcmDataExports_tagResourceCmd.MarkFlagRequired("resource-arn")
		bcmDataExports_tagResourceCmd.MarkFlagRequired("resource-tags")
	})
	bcmDataExportsCmd.AddCommand(bcmDataExports_tagResourceCmd)
}
