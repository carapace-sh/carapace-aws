package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingsCmd = &cobra.Command{
	Use:   "list-things",
	Short: "Lists your things.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listThingsCmd).Standalone()

		iot_listThingsCmd.Flags().String("attribute-name", "", "The attribute name used to search for things.")
		iot_listThingsCmd.Flags().String("attribute-value", "", "The attribute value used to search for things.")
		iot_listThingsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iot_listThingsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listThingsCmd.Flags().String("thing-type-name", "", "The name of the thing type used to search for things.")
		iot_listThingsCmd.Flags().String("use-prefix-attribute-value", "", "When `true`, the action returns the thing resources with attribute values that start with the `attributeValue` provided.")
	})
	iotCmd.AddCommand(iot_listThingsCmd)
}
