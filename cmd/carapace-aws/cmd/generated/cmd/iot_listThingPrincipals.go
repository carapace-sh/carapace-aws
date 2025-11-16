package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingPrincipalsCmd = &cobra.Command{
	Use:   "list-thing-principals",
	Short: "Lists the principals associated with the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingPrincipalsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listThingPrincipalsCmd).Standalone()

		iot_listThingPrincipalsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iot_listThingPrincipalsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listThingPrincipalsCmd.Flags().String("thing-name", "", "The name of the thing.")
		iot_listThingPrincipalsCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_listThingPrincipalsCmd)
}
