package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getEntitiesCmd = &cobra.Command{
	Use:   "get-entities",
	Short: "Gets definitions of the specified entities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_getEntitiesCmd).Standalone()

		iotthingsgraph_getEntitiesCmd.Flags().String("ids", "", "An array of entity IDs.")
		iotthingsgraph_getEntitiesCmd.Flags().String("namespace-version", "", "The version of the user's namespace.")
		iotthingsgraph_getEntitiesCmd.MarkFlagRequired("ids")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getEntitiesCmd)
}
