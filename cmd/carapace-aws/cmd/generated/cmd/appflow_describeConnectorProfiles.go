package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeConnectorProfilesCmd = &cobra.Command{
	Use:   "describe-connector-profiles",
	Short: "Returns a list of `connector-profile` details matching the provided `connector-profile` names and `connector-types`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeConnectorProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_describeConnectorProfilesCmd).Standalone()

		appflow_describeConnectorProfilesCmd.Flags().String("connector-label", "", "The name of the connector.")
		appflow_describeConnectorProfilesCmd.Flags().String("connector-profile-names", "", "The name of the connector profile.")
		appflow_describeConnectorProfilesCmd.Flags().String("connector-type", "", "The type of connector, such as Salesforce, Amplitude, and so on.")
		appflow_describeConnectorProfilesCmd.Flags().String("max-results", "", "Specifies the maximum number of items that should be returned in the result set.")
		appflow_describeConnectorProfilesCmd.Flags().String("next-token", "", "The pagination token for the next page of data.")
	})
	appflowCmd.AddCommand(appflow_describeConnectorProfilesCmd)
}
