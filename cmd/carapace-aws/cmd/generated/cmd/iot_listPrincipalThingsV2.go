package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPrincipalThingsV2Cmd = &cobra.Command{
	Use:   "list-principal-things-v2",
	Short: "Lists the things associated with the specified principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPrincipalThingsV2Cmd).Standalone()

	iot_listPrincipalThingsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listPrincipalThingsV2Cmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listPrincipalThingsV2Cmd.Flags().String("principal", "", "The principal.")
	iot_listPrincipalThingsV2Cmd.Flags().String("thing-principal-type", "", "The type of the relation you want to filter in the response.")
	iot_listPrincipalThingsV2Cmd.MarkFlagRequired("principal")
	iotCmd.AddCommand(iot_listPrincipalThingsV2Cmd)
}
