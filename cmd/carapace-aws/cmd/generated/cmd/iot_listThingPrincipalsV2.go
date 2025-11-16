package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingPrincipalsV2Cmd = &cobra.Command{
	Use:   "list-thing-principals-v2",
	Short: "Lists the principals associated with the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingPrincipalsV2Cmd).Standalone()

	iot_listThingPrincipalsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listThingPrincipalsV2Cmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listThingPrincipalsV2Cmd.Flags().String("thing-name", "", "The name of the thing.")
	iot_listThingPrincipalsV2Cmd.Flags().String("thing-principal-type", "", "The type of the relation you want to filter in the response.")
	iot_listThingPrincipalsV2Cmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_listThingPrincipalsV2Cmd)
}
