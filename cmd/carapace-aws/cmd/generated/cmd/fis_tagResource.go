package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_tagResourceCmd).Standalone()

	fis_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	fis_tagResourceCmd.Flags().String("tags", "", "The tags for the resource.")
	fis_tagResourceCmd.MarkFlagRequired("resource-arn")
	fis_tagResourceCmd.MarkFlagRequired("tags")
	fisCmd.AddCommand(fis_tagResourceCmd)
}
