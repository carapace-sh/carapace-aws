package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Gets an Amazon DataZone data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getDataSourceCmd).Standalone()

	datazone_getDataSourceCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the data source exists.")
	datazone_getDataSourceCmd.Flags().String("identifier", "", "The ID of the Amazon DataZone data source.")
	datazone_getDataSourceCmd.MarkFlagRequired("domain-identifier")
	datazone_getDataSourceCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getDataSourceCmd)
}
