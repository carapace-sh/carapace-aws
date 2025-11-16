package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_deleteDataLakeDatasetCmd = &cobra.Command{
	Use:   "delete-data-lake-dataset",
	Short: "Enables you to programmatically delete an Amazon Web Services Supply Chain data lake dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_deleteDataLakeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_deleteDataLakeDatasetCmd).Standalone()

		supplychain_deleteDataLakeDatasetCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
		supplychain_deleteDataLakeDatasetCmd.Flags().String("name", "", "The name of the dataset.")
		supplychain_deleteDataLakeDatasetCmd.Flags().String("namespace", "", "The namespace of the dataset, besides the custom defined namespace, every instance comes with below pre-defined namespaces:")
		supplychain_deleteDataLakeDatasetCmd.MarkFlagRequired("instance-id")
		supplychain_deleteDataLakeDatasetCmd.MarkFlagRequired("name")
		supplychain_deleteDataLakeDatasetCmd.MarkFlagRequired("namespace")
	})
	supplychainCmd.AddCommand(supplychain_deleteDataLakeDatasetCmd)
}
