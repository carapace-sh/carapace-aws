package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeConnectorCmd = &cobra.Command{
	Use:   "describe-connector",
	Short: "Describes the given custom connector registered in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_describeConnectorCmd).Standalone()

		appflow_describeConnectorCmd.Flags().String("connector-label", "", "The label of the connector.")
		appflow_describeConnectorCmd.Flags().String("connector-type", "", "The connector type, such as CUSTOMCONNECTOR, Saleforce, Marketo.")
		appflow_describeConnectorCmd.MarkFlagRequired("connector-type")
	})
	appflowCmd.AddCommand(appflow_describeConnectorCmd)
}
