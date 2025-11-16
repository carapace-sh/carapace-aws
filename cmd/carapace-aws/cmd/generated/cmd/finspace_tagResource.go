package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds metadata tags to a FinSpace resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_tagResourceCmd).Standalone()

		finspace_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource.")
		finspace_tagResourceCmd.Flags().String("tags", "", "One or more tags to be assigned to the resource.")
		finspace_tagResourceCmd.MarkFlagRequired("resource-arn")
		finspace_tagResourceCmd.MarkFlagRequired("tags")
	})
	finspaceCmd.AddCommand(finspace_tagResourceCmd)
}
