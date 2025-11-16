package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_createDatastoreCmd = &cobra.Command{
	Use:   "create-datastore",
	Short: "Creates a data store, which is a repository for messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_createDatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_createDatastoreCmd).Standalone()

		iotanalytics_createDatastoreCmd.Flags().String("datastore-name", "", "The name of the data store.")
		iotanalytics_createDatastoreCmd.Flags().String("datastore-partitions", "", "Contains information about the partition dimensions in a data store.")
		iotanalytics_createDatastoreCmd.Flags().String("datastore-storage", "", "Where data in a data store is stored.. You can choose `serviceManagedS3` storage, `customerManagedS3` storage, or `iotSiteWiseMultiLayerStorage` storage.")
		iotanalytics_createDatastoreCmd.Flags().String("file-format-configuration", "", "Contains the configuration information of file formats.")
		iotanalytics_createDatastoreCmd.Flags().String("retention-period", "", "How long, in days, message data is kept for the data store.")
		iotanalytics_createDatastoreCmd.Flags().String("tags", "", "Metadata which can be used to manage the data store.")
		iotanalytics_createDatastoreCmd.MarkFlagRequired("datastore-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_createDatastoreCmd)
}
