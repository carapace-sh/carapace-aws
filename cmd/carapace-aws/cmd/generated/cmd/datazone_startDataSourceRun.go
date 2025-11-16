package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_startDataSourceRunCmd = &cobra.Command{
	Use:   "start-data-source-run",
	Short: "Start the run of the specified data source in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_startDataSourceRunCmd).Standalone()

	datazone_startDataSourceRunCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_startDataSourceRunCmd.Flags().String("data-source-identifier", "", "The identifier of the data source.")
	datazone_startDataSourceRunCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to start a data source run.")
	datazone_startDataSourceRunCmd.MarkFlagRequired("data-source-identifier")
	datazone_startDataSourceRunCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_startDataSourceRunCmd)
}
