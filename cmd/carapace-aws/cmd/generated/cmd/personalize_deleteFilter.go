package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteFilterCmd = &cobra.Command{
	Use:   "delete-filter",
	Short: "Deletes a filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteFilterCmd).Standalone()

	personalize_deleteFilterCmd.Flags().String("filter-arn", "", "The ARN of the filter to delete.")
	personalize_deleteFilterCmd.MarkFlagRequired("filter-arn")
	personalizeCmd.AddCommand(personalize_deleteFilterCmd)
}
