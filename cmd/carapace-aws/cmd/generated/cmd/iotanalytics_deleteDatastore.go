package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_deleteDatastoreCmd = &cobra.Command{
	Use:   "delete-datastore",
	Short: "Deletes the specified data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_deleteDatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_deleteDatastoreCmd).Standalone()

		iotanalytics_deleteDatastoreCmd.Flags().String("datastore-name", "", "The name of the data store to delete.")
		iotanalytics_deleteDatastoreCmd.MarkFlagRequired("datastore-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_deleteDatastoreCmd)
}
