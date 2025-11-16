package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a medical imaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_untagResourceCmd).Standalone()

	medicalImaging_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the medical imaging resource that tags are being removed from.")
	medicalImaging_untagResourceCmd.Flags().String("tag-keys", "", "The keys for the tags to be removed from the medical imaging resource.")
	medicalImaging_untagResourceCmd.MarkFlagRequired("resource-arn")
	medicalImaging_untagResourceCmd.MarkFlagRequired("tag-keys")
	medicalImagingCmd.AddCommand(medicalImaging_untagResourceCmd)
}
