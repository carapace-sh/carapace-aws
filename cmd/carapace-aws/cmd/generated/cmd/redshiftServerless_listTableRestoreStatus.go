package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listTableRestoreStatusCmd = &cobra.Command{
	Use:   "list-table-restore-status",
	Short: "Returns information about an array of `TableRestoreStatus` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listTableRestoreStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listTableRestoreStatusCmd).Standalone()

		redshiftServerless_listTableRestoreStatusCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listTableRestoreStatusCmd.Flags().String("namespace-name", "", "The namespace from which to list all of the statuses of `RestoreTableFromSnapshot` operations .")
		redshiftServerless_listTableRestoreStatusCmd.Flags().String("next-token", "", "If your initial `ListTableRestoreStatus` operation returns a nextToken, you can include the returned `nextToken` in following `ListTableRestoreStatus` operations.")
		redshiftServerless_listTableRestoreStatusCmd.Flags().String("workgroup-name", "", "The workgroup from which to list all of the statuses of `RestoreTableFromSnapshot` operations.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listTableRestoreStatusCmd)
}
