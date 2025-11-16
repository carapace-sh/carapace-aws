package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getFindingsFilterCmd = &cobra.Command{
	Use:   "get-findings-filter",
	Short: "Retrieves the criteria and other settings for a findings filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getFindingsFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getFindingsFilterCmd).Standalone()

		macie2_getFindingsFilterCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_getFindingsFilterCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_getFindingsFilterCmd)
}
