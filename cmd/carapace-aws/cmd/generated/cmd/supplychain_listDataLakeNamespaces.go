package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listDataLakeNamespacesCmd = &cobra.Command{
	Use:   "list-data-lake-namespaces",
	Short: "Enables you to programmatically view the list of Amazon Web Services Supply Chain data lake namespaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listDataLakeNamespacesCmd).Standalone()

	supplychain_listDataLakeNamespacesCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_listDataLakeNamespacesCmd.Flags().String("max-results", "", "The max number of namespaces to fetch in this paginated request.")
	supplychain_listDataLakeNamespacesCmd.Flags().String("next-token", "", "The pagination token to fetch next page of namespaces.")
	supplychain_listDataLakeNamespacesCmd.MarkFlagRequired("instance-id")
	supplychainCmd.AddCommand(supplychain_listDataLakeNamespacesCmd)
}
