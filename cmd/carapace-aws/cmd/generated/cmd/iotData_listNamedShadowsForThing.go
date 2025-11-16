package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_listNamedShadowsForThingCmd = &cobra.Command{
	Use:   "list-named-shadows-for-thing",
	Short: "Lists the shadows for the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_listNamedShadowsForThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotData_listNamedShadowsForThingCmd).Standalone()

		iotData_listNamedShadowsForThingCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
		iotData_listNamedShadowsForThingCmd.Flags().String("page-size", "", "The result page size.")
		iotData_listNamedShadowsForThingCmd.Flags().String("thing-name", "", "The name of the thing.")
		iotData_listNamedShadowsForThingCmd.MarkFlagRequired("thing-name")
	})
	iotDataCmd.AddCommand(iotData_listNamedShadowsForThingCmd)
}
