package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_writeRecordsCmd = &cobra.Command{
	Use:   "write-records",
	Short: "Enables you to write your time-series data into Timestream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_writeRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_writeRecordsCmd).Standalone()

		timestreamWrite_writeRecordsCmd.Flags().String("common-attributes", "", "A record that contains the common measure, dimension, time, and version attributes shared across all the records in the request.")
		timestreamWrite_writeRecordsCmd.Flags().String("database-name", "", "The name of the Timestream database.")
		timestreamWrite_writeRecordsCmd.Flags().String("records", "", "An array of records that contain the unique measure, dimension, time, and version attributes for each time-series data point.")
		timestreamWrite_writeRecordsCmd.Flags().String("table-name", "", "The name of the Timestream table.")
		timestreamWrite_writeRecordsCmd.MarkFlagRequired("database-name")
		timestreamWrite_writeRecordsCmd.MarkFlagRequired("records")
		timestreamWrite_writeRecordsCmd.MarkFlagRequired("table-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_writeRecordsCmd)
}
