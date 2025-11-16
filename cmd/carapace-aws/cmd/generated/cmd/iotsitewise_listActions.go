package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listActionsCmd = &cobra.Command{
	Use:   "list-actions",
	Short: "Retrieves a paginated list of actions for a specific target resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listActionsCmd).Standalone()

		iotsitewise_listActionsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listActionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listActionsCmd.Flags().String("resolve-to-resource-id", "", "The ID of the resolved resource.")
		iotsitewise_listActionsCmd.Flags().String("resolve-to-resource-type", "", "The type of the resolved resource.")
		iotsitewise_listActionsCmd.Flags().String("target-resource-id", "", "The ID of the target resource.")
		iotsitewise_listActionsCmd.Flags().String("target-resource-type", "", "The type of resource.")
		iotsitewise_listActionsCmd.MarkFlagRequired("target-resource-id")
		iotsitewise_listActionsCmd.MarkFlagRequired("target-resource-type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listActionsCmd)
}
