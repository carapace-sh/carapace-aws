package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_listConnectorEntitiesCmd = &cobra.Command{
	Use:   "list-connector-entities",
	Short: "Returns the list of available connector entities supported by Amazon AppFlow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_listConnectorEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_listConnectorEntitiesCmd).Standalone()

		appflow_listConnectorEntitiesCmd.Flags().String("api-version", "", "The version of the API that's used by the connector.")
		appflow_listConnectorEntitiesCmd.Flags().String("connector-profile-name", "", "The name of the connector profile.")
		appflow_listConnectorEntitiesCmd.Flags().String("connector-type", "", "The type of connector, such as Salesforce, Amplitude, and so on.")
		appflow_listConnectorEntitiesCmd.Flags().String("entities-path", "", "This optional parameter is specific to connector implementation.")
		appflow_listConnectorEntitiesCmd.Flags().String("max-results", "", "The maximum number of items that the operation returns in the response.")
		appflow_listConnectorEntitiesCmd.Flags().String("next-token", "", "A token that was provided by your prior `ListConnectorEntities` operation if the response was too big for the page size.")
	})
	appflowCmd.AddCommand(appflow_listConnectorEntitiesCmd)
}
