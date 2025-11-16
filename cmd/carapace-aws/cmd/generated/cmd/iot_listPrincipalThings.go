package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPrincipalThingsCmd = &cobra.Command{
	Use:   "list-principal-things",
	Short: "Lists the things associated with the specified principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPrincipalThingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listPrincipalThingsCmd).Standalone()

		iot_listPrincipalThingsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iot_listPrincipalThingsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listPrincipalThingsCmd.Flags().String("principal", "", "The principal.")
		iot_listPrincipalThingsCmd.MarkFlagRequired("principal")
	})
	iotCmd.AddCommand(iot_listPrincipalThingsCmd)
}
