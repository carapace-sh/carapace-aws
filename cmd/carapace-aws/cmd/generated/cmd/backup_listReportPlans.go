package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listReportPlansCmd = &cobra.Command{
	Use:   "list-report-plans",
	Short: "Returns a list of your report plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listReportPlansCmd).Standalone()

	backup_listReportPlansCmd.Flags().String("max-results", "", "The number of desired results from 1 to 1000. Optional.")
	backup_listReportPlansCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	backupCmd.AddCommand(backup_listReportPlansCmd)
}
