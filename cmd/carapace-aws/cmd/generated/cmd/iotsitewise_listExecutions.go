package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listExecutionsCmd = &cobra.Command{
	Use:   "list-executions",
	Short: "Retrieves a paginated list of summaries of all executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listExecutionsCmd).Standalone()

		iotsitewise_listExecutionsCmd.Flags().String("action-type", "", "The type of action exectued.")
		iotsitewise_listExecutionsCmd.Flags().String("max-results", "", "The maximum number of results returned for each paginated request.")
		iotsitewise_listExecutionsCmd.Flags().String("next-token", "", "The token used for the next set of paginated results.")
		iotsitewise_listExecutionsCmd.Flags().String("resolve-to-resource-id", "", "The ID of the resolved resource.")
		iotsitewise_listExecutionsCmd.Flags().String("resolve-to-resource-type", "", "The type of the resolved resource.")
		iotsitewise_listExecutionsCmd.Flags().String("target-resource-id", "", "The ID of the target resource.")
		iotsitewise_listExecutionsCmd.Flags().String("target-resource-type", "", "The type of the target resource.")
		iotsitewise_listExecutionsCmd.MarkFlagRequired("target-resource-id")
		iotsitewise_listExecutionsCmd.MarkFlagRequired("target-resource-type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listExecutionsCmd)
}
