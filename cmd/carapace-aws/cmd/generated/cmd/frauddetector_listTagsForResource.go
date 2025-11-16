package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_listTagsForResourceCmd).Standalone()

	frauddetector_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
	frauddetector_listTagsForResourceCmd.Flags().String("next-token", "", "The next token from the previous results.")
	frauddetector_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN that specifies the resource whose tags you want to list.")
	frauddetector_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	frauddetectorCmd.AddCommand(frauddetector_listTagsForResourceCmd)
}
