package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_tagResourceCmd).Standalone()

		omics_tagResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
		omics_tagResourceCmd.Flags().String("tags", "", "Tags for the resource.")
		omics_tagResourceCmd.MarkFlagRequired("resource-arn")
		omics_tagResourceCmd.MarkFlagRequired("tags")
	})
	omicsCmd.AddCommand(omics_tagResourceCmd)
}
