package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listDatastoresCmd = &cobra.Command{
	Use:   "list-datastores",
	Short: "Retrieves a list of data stores.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listDatastoresCmd).Standalone()

	iotanalytics_listDatastoresCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	iotanalytics_listDatastoresCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iotanalyticsCmd.AddCommand(iotanalytics_listDatastoresCmd)
}
