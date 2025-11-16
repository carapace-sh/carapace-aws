package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_deleteFindingsFilterCmd = &cobra.Command{
	Use:   "delete-findings-filter",
	Short: "Deletes a findings filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_deleteFindingsFilterCmd).Standalone()

	macie2_deleteFindingsFilterCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_deleteFindingsFilterCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_deleteFindingsFilterCmd)
}
