package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listComputationModelResolveToResourcesCmd = &cobra.Command{
	Use:   "list-computation-model-resolve-to-resources",
	Short: "Lists all distinct resources that are resolved from the executed actions of the computation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listComputationModelResolveToResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listComputationModelResolveToResourcesCmd).Standalone()

		iotsitewise_listComputationModelResolveToResourcesCmd.Flags().String("computation-model-id", "", "The ID of the computation model for which to list resolved resources.")
		iotsitewise_listComputationModelResolveToResourcesCmd.Flags().String("max-results", "", "The maximum number of results returned for each paginated request.")
		iotsitewise_listComputationModelResolveToResourcesCmd.Flags().String("next-token", "", "The token used for the next set of paginated results.")
		iotsitewise_listComputationModelResolveToResourcesCmd.MarkFlagRequired("computation-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listComputationModelResolveToResourcesCmd)
}
