package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_createDataLakeDatasetCmd = &cobra.Command{
	Use:   "create-data-lake-dataset",
	Short: "Enables you to programmatically create an Amazon Web Services Supply Chain data lake dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_createDataLakeDatasetCmd).Standalone()

	supplychain_createDataLakeDatasetCmd.Flags().String("description", "", "The description of the dataset.")
	supplychain_createDataLakeDatasetCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_createDataLakeDatasetCmd.Flags().String("name", "", "The name of the dataset.")
	supplychain_createDataLakeDatasetCmd.Flags().String("namespace", "", "The namespace of the dataset, besides the custom defined namespace, every instance comes with below pre-defined namespaces:")
	supplychain_createDataLakeDatasetCmd.Flags().String("partition-spec", "", "The partition specification of the dataset.")
	supplychain_createDataLakeDatasetCmd.Flags().String("schema", "", "The custom schema of the data lake dataset and required for dataset in **default** and custom namespaces.")
	supplychain_createDataLakeDatasetCmd.Flags().String("tags", "", "The tags of the dataset.")
	supplychain_createDataLakeDatasetCmd.MarkFlagRequired("instance-id")
	supplychain_createDataLakeDatasetCmd.MarkFlagRequired("name")
	supplychain_createDataLakeDatasetCmd.MarkFlagRequired("namespace")
	supplychainCmd.AddCommand(supplychain_createDataLakeDatasetCmd)
}
