package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteDataProductCmd = &cobra.Command{
	Use:   "delete-data-product",
	Short: "Deletes a data product in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteDataProductCmd).Standalone()

	datazone_deleteDataProductCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the data product is deleted.")
	datazone_deleteDataProductCmd.Flags().String("identifier", "", "The identifier of the data product that is deleted.")
	datazone_deleteDataProductCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteDataProductCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteDataProductCmd)
}
