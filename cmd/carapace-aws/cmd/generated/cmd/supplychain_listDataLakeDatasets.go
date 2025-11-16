package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listDataLakeDatasetsCmd = &cobra.Command{
	Use:   "list-data-lake-datasets",
	Short: "Enables you to programmatically view the list of Amazon Web Services Supply Chain data lake datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listDataLakeDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_listDataLakeDatasetsCmd).Standalone()

		supplychain_listDataLakeDatasetsCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_listDataLakeDatasetsCmd.Flags().String("max-results", "", "The max number of datasets to fetch in this paginated request.")
		supplychain_listDataLakeDatasetsCmd.Flags().String("namespace", "", "The namespace of the dataset, besides the custom defined namespace, every instance comes with below pre-defined namespaces:")
		supplychain_listDataLakeDatasetsCmd.Flags().String("next-token", "", "The pagination token to fetch next page of datasets.")
		supplychain_listDataLakeDatasetsCmd.MarkFlagRequired("instance-id")
		supplychain_listDataLakeDatasetsCmd.MarkFlagRequired("namespace")
	})
	supplychainCmd.AddCommand(supplychain_listDataLakeDatasetsCmd)
}
