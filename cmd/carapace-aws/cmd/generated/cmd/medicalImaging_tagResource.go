package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a user-specifed key and value tag to a medical imaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_tagResourceCmd).Standalone()

	medicalImaging_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the medical imaging resource that tags are being added to.")
	medicalImaging_tagResourceCmd.Flags().String("tags", "", "The user-specified key and value tag pairs added to a medical imaging resource.")
	medicalImaging_tagResourceCmd.MarkFlagRequired("resource-arn")
	medicalImaging_tagResourceCmd.MarkFlagRequired("tags")
	medicalImagingCmd.AddCommand(medicalImaging_tagResourceCmd)
}
