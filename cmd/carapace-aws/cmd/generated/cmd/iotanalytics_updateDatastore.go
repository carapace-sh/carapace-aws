package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_updateDatastoreCmd = &cobra.Command{
	Use:   "update-datastore",
	Short: "Used to update the settings of a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_updateDatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_updateDatastoreCmd).Standalone()

		iotanalytics_updateDatastoreCmd.Flags().String("datastore-name", "", "The name of the data store to be updated.")
		iotanalytics_updateDatastoreCmd.Flags().String("datastore-storage", "", "Where data in a data store is stored.. You can choose `serviceManagedS3` storage, `customerManagedS3` storage, or `iotSiteWiseMultiLayerStorage` storage.")
		iotanalytics_updateDatastoreCmd.Flags().String("file-format-configuration", "", "Contains the configuration information of file formats.")
		iotanalytics_updateDatastoreCmd.Flags().String("retention-period", "", "How long, in days, message data is kept for the data store.")
		iotanalytics_updateDatastoreCmd.MarkFlagRequired("datastore-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_updateDatastoreCmd)
}
