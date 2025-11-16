package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dlm_tagResourceCmd).Standalone()

		dlm_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		dlm_tagResourceCmd.Flags().String("tags", "", "One or more tags.")
		dlm_tagResourceCmd.MarkFlagRequired("resource-arn")
		dlm_tagResourceCmd.MarkFlagRequired("tags")
	})
	dlmCmd.AddCommand(dlm_tagResourceCmd)
}
