package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listIndicesCmd = &cobra.Command{
	Use:   "list-indices",
	Short: "Lists the search indices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listIndicesCmd).Standalone()

	iot_listIndicesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listIndicesCmd.Flags().String("next-token", "", "The token used to get the next set of results, or `null` if there are no additional results.")
	iotCmd.AddCommand(iot_listIndicesCmd)
}
