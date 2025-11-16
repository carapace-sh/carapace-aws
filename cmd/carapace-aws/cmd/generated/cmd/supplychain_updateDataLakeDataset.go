package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_updateDataLakeDatasetCmd = &cobra.Command{
	Use:   "update-data-lake-dataset",
	Short: "Enables you to programmatically update an Amazon Web Services Supply Chain data lake dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_updateDataLakeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_updateDataLakeDatasetCmd).Standalone()

		supplychain_updateDataLakeDatasetCmd.Flags().String("description", "", "The updated description of the data lake dataset.")
		supplychain_updateDataLakeDatasetCmd.Flags().String("instance-id", "", "The Amazon Web Services Chain instance identifier.")
		supplychain_updateDataLakeDatasetCmd.Flags().String("name", "", "The name of the dataset.")
		supplychain_updateDataLakeDatasetCmd.Flags().String("namespace", "", "The namespace of the dataset, besides the custom defined namespace, every instance comes with below pre-defined namespaces:")
		supplychain_updateDataLakeDatasetCmd.MarkFlagRequired("instance-id")
		supplychain_updateDataLakeDatasetCmd.MarkFlagRequired("name")
		supplychain_updateDataLakeDatasetCmd.MarkFlagRequired("namespace")
	})
	supplychainCmd.AddCommand(supplychain_updateDataLakeDatasetCmd)
}
