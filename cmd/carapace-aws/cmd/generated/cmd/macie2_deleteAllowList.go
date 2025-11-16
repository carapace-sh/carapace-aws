package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_deleteAllowListCmd = &cobra.Command{
	Use:   "delete-allow-list",
	Short: "Deletes an allow list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_deleteAllowListCmd).Standalone()

	macie2_deleteAllowListCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_deleteAllowListCmd.Flags().String("ignore-job-checks", "", "Specifies whether to force deletion of the allow list, even if active classification jobs are configured to use the list.")
	macie2_deleteAllowListCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_deleteAllowListCmd)
}
