package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listComputationModelDataBindingUsagesCmd = &cobra.Command{
	Use:   "list-computation-model-data-binding-usages",
	Short: "Lists all data binding usages for computation models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listComputationModelDataBindingUsagesCmd).Standalone()

	iotsitewise_listComputationModelDataBindingUsagesCmd.Flags().String("data-binding-value-filter", "", "A filter used to limit the returned data binding usages based on specific data binding values.")
	iotsitewise_listComputationModelDataBindingUsagesCmd.Flags().String("max-results", "", "The maximum number of results returned for each paginated request.")
	iotsitewise_listComputationModelDataBindingUsagesCmd.Flags().String("next-token", "", "The token used for the next set of paginated results.")
	iotsitewise_listComputationModelDataBindingUsagesCmd.MarkFlagRequired("data-binding-value-filter")
	iotsitewiseCmd.AddCommand(iotsitewise_listComputationModelDataBindingUsagesCmd)
}
