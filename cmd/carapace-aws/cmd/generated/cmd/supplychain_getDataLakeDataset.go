package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getDataLakeDatasetCmd = &cobra.Command{
	Use:   "get-data-lake-dataset",
	Short: "Enables you to programmatically view an Amazon Web Services Supply Chain data lake dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getDataLakeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_getDataLakeDatasetCmd).Standalone()

		supplychain_getDataLakeDatasetCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_getDataLakeDatasetCmd.Flags().String("name", "", "The name of the dataset.")
		supplychain_getDataLakeDatasetCmd.Flags().String("namespace", "", "The namespace of the dataset, besides the custom defined namespace, every instance comes with below pre-defined namespaces:")
		supplychain_getDataLakeDatasetCmd.MarkFlagRequired("instance-id")
		supplychain_getDataLakeDatasetCmd.MarkFlagRequired("name")
		supplychain_getDataLakeDatasetCmd.MarkFlagRequired("namespace")
	})
	supplychainCmd.AddCommand(supplychain_getDataLakeDatasetCmd)
}
