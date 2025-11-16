package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeBackupsCmd = &cobra.Command{
	Use:   "describe-backups",
	Short: "Returns the description of a specific Amazon FSx backup, if a `BackupIds` value is provided for that backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeBackupsCmd).Standalone()

	fsx_describeBackupsCmd.Flags().String("backup-ids", "", "The IDs of the backups that you want to retrieve.")
	fsx_describeBackupsCmd.Flags().String("filters", "", "The filters structure.")
	fsx_describeBackupsCmd.Flags().String("max-results", "", "Maximum number of backups to return in the response.")
	fsx_describeBackupsCmd.Flags().String("next-token", "", "An opaque pagination token returned from a previous `DescribeBackups` operation.")
	fsxCmd.AddCommand(fsx_describeBackupsCmd)
}
