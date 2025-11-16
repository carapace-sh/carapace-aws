package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a medical imaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_listTagsForResourceCmd).Standalone()

		medicalImaging_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the medical imaging resource to list tags for.")
		medicalImaging_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	medicalImagingCmd.AddCommand(medicalImaging_listTagsForResourceCmd)
}
