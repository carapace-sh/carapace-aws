package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listIndexedRecoveryPointsCmd = &cobra.Command{
	Use:   "list-indexed-recovery-points",
	Short: "This operation returns a list of recovery points that have an associated index, belonging to the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listIndexedRecoveryPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listIndexedRecoveryPointsCmd).Standalone()

		backup_listIndexedRecoveryPointsCmd.Flags().String("created-after", "", "Returns only indexed recovery points that were created after the specified date.")
		backup_listIndexedRecoveryPointsCmd.Flags().String("created-before", "", "Returns only indexed recovery points that were created before the specified date.")
		backup_listIndexedRecoveryPointsCmd.Flags().String("index-status", "", "Include this parameter to filter the returned list by the indicated statuses.")
		backup_listIndexedRecoveryPointsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
		backup_listIndexedRecoveryPointsCmd.Flags().String("next-token", "", "The next item following a partial list of returned recovery points.")
		backup_listIndexedRecoveryPointsCmd.Flags().String("resource-type", "", "Returns a list of indexed recovery points for the specified resource type(s).")
		backup_listIndexedRecoveryPointsCmd.Flags().String("source-resource-arn", "", "A string of the Amazon Resource Name (ARN) that uniquely identifies the source resource.")
	})
	backupCmd.AddCommand(backup_listIndexedRecoveryPointsCmd)
}
