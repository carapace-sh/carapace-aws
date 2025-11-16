package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Lists all datasets currently available in your account, filtering on the dataset name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listDatasetsCmd).Standalone()

	lookoutequipment_listDatasetsCmd.Flags().String("dataset-name-begins-with", "", "The beginning of the name of the datasets to be listed.")
	lookoutequipment_listDatasetsCmd.Flags().String("max-results", "", "Specifies the maximum number of datasets to list.")
	lookoutequipment_listDatasetsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of datasets.")
	lookoutequipmentCmd.AddCommand(lookoutequipment_listDatasetsCmd)
}
