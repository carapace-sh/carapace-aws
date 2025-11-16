package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDblogFilesCmd = &cobra.Command{
	Use:   "describe-dblog-files",
	Short: "Returns a list of DB log files for the DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDblogFilesCmd).Standalone()

	rds_describeDblogFilesCmd.Flags().String("dbinstance-identifier", "", "The customer-assigned name of the DB instance that contains the log files you want to list.")
	rds_describeDblogFilesCmd.Flags().String("file-last-written", "", "Filters the available log files for files written since the specified date, in POSIX timestamp format with milliseconds.")
	rds_describeDblogFilesCmd.Flags().String("file-size", "", "Filters the available log files for files larger than the specified size.")
	rds_describeDblogFilesCmd.Flags().String("filename-contains", "", "Filters the available log files for log file names that contain the specified string.")
	rds_describeDblogFilesCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeDblogFilesCmd.Flags().String("marker", "", "The pagination token provided in the previous request.")
	rds_describeDblogFilesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDblogFilesCmd.MarkFlagRequired("dbinstance-identifier")
	rdsCmd.AddCommand(rds_describeDblogFilesCmd)
}
