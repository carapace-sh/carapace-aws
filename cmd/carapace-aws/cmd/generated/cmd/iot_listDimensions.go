package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listDimensionsCmd = &cobra.Command{
	Use:   "list-dimensions",
	Short: "List the set of dimensions that are defined for your Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listDimensionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listDimensionsCmd).Standalone()

		iot_listDimensionsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve at one time.")
		iot_listDimensionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	iotCmd.AddCommand(iot_listDimensionsCmd)
}
