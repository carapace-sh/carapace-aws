package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or edits tags on an PCS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_tagResourceCmd).Standalone()

	pcs_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pcs_tagResourceCmd.Flags().String("tags", "", "1 or more tags added to the resource.")
	pcs_tagResourceCmd.MarkFlagRequired("resource-arn")
	pcs_tagResourceCmd.MarkFlagRequired("tags")
	pcsCmd.AddCommand(pcs_tagResourceCmd)
}
