package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingTypesCmd = &cobra.Command{
	Use:   "list-thing-types",
	Short: "Lists the existing thing types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingTypesCmd).Standalone()

	iot_listThingTypesCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listThingTypesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listThingTypesCmd.Flags().String("thing-type-name", "", "The name of the thing type.")
	iotCmd.AddCommand(iot_listThingTypesCmd)
}
