package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getDataProductCmd = &cobra.Command{
	Use:   "get-data-product",
	Short: "Gets the data product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getDataProductCmd).Standalone()

	datazone_getDataProductCmd.Flags().String("domain-identifier", "", "The ID of the domain where the data product lives.")
	datazone_getDataProductCmd.Flags().String("identifier", "", "The ID of the data product.")
	datazone_getDataProductCmd.Flags().String("revision", "", "The revision of the data product.")
	datazone_getDataProductCmd.MarkFlagRequired("domain-identifier")
	datazone_getDataProductCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getDataProductCmd)
}
