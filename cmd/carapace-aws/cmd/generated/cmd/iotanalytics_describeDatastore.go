package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_describeDatastoreCmd = &cobra.Command{
	Use:   "describe-datastore",
	Short: "Retrieves information about a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_describeDatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_describeDatastoreCmd).Standalone()

		iotanalytics_describeDatastoreCmd.Flags().String("datastore-name", "", "The name of the data store")
		iotanalytics_describeDatastoreCmd.Flags().String("include-statistics", "", "If true, additional statistical information about the data store is included in the response.")
		iotanalytics_describeDatastoreCmd.MarkFlagRequired("datastore-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_describeDatastoreCmd)
}
