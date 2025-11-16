package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeConnectorEntityCmd = &cobra.Command{
	Use:   "describe-connector-entity",
	Short: "Provides details regarding the entity used with the connector, with a description of the data model for each field in that entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeConnectorEntityCmd).Standalone()

	appflow_describeConnectorEntityCmd.Flags().String("api-version", "", "The version of the API that's used by the connector.")
	appflow_describeConnectorEntityCmd.Flags().String("connector-entity-name", "", "The entity name for that connector.")
	appflow_describeConnectorEntityCmd.Flags().String("connector-profile-name", "", "The name of the connector profile.")
	appflow_describeConnectorEntityCmd.Flags().String("connector-type", "", "The type of connector application, such as Salesforce, Amplitude, and so on.")
	appflow_describeConnectorEntityCmd.MarkFlagRequired("connector-entity-name")
	appflowCmd.AddCommand(appflow_describeConnectorEntityCmd)
}
