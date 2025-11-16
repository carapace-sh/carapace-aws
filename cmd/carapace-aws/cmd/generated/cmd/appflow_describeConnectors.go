package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeConnectorsCmd = &cobra.Command{
	Use:   "describe-connectors",
	Short: "Describes the connectors vended by Amazon AppFlow for specified connector types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeConnectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_describeConnectorsCmd).Standalone()

		appflow_describeConnectorsCmd.Flags().String("connector-types", "", "The type of connector, such as Salesforce, Amplitude, and so on.")
		appflow_describeConnectorsCmd.Flags().String("max-results", "", "The maximum number of items that should be returned in the result set.")
		appflow_describeConnectorsCmd.Flags().String("next-token", "", "The pagination token for the next page of data.")
	})
	appflowCmd.AddCommand(appflow_describeConnectorsCmd)
}
