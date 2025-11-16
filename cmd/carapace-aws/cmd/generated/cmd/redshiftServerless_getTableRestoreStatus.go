package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getTableRestoreStatusCmd = &cobra.Command{
	Use:   "get-table-restore-status",
	Short: "Returns information about a `TableRestoreStatus` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getTableRestoreStatusCmd).Standalone()

	redshiftServerless_getTableRestoreStatusCmd.Flags().String("table-restore-request-id", "", "The ID of the `RestoreTableFromSnapshot` request to return status for.")
	redshiftServerless_getTableRestoreStatusCmd.MarkFlagRequired("table-restore-request-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getTableRestoreStatusCmd)
}
