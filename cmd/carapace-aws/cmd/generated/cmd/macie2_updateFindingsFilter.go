package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateFindingsFilterCmd = &cobra.Command{
	Use:   "update-findings-filter",
	Short: "Updates the criteria and other settings for a findings filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateFindingsFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateFindingsFilterCmd).Standalone()

		macie2_updateFindingsFilterCmd.Flags().String("action", "", "The action to perform on findings that match the filter criteria (findingCriteria).")
		macie2_updateFindingsFilterCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		macie2_updateFindingsFilterCmd.Flags().String("description", "", "A custom description of the filter.")
		macie2_updateFindingsFilterCmd.Flags().String("finding-criteria", "", "The criteria to use to filter findings.")
		macie2_updateFindingsFilterCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_updateFindingsFilterCmd.Flags().String("name", "", "A custom name for the filter.")
		macie2_updateFindingsFilterCmd.Flags().String("position", "", "The position of the filter in the list of saved filters on the Amazon Macie console.")
		macie2_updateFindingsFilterCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_updateFindingsFilterCmd)
}
