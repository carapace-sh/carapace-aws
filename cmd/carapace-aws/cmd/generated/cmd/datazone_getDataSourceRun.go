package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getDataSourceRunCmd = &cobra.Command{
	Use:   "get-data-source-run",
	Short: "Gets an Amazon DataZone data source run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getDataSourceRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getDataSourceRunCmd).Standalone()

		datazone_getDataSourceRunCmd.Flags().String("domain-identifier", "", "The ID of the domain in which this data source run was performed.")
		datazone_getDataSourceRunCmd.Flags().String("identifier", "", "The ID of the data source run.")
		datazone_getDataSourceRunCmd.MarkFlagRequired("domain-identifier")
		datazone_getDataSourceRunCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getDataSourceRunCmd)
}
