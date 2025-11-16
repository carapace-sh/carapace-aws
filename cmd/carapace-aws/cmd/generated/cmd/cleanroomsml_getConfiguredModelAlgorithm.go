package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getConfiguredModelAlgorithmCmd = &cobra.Command{
	Use:   "get-configured-model-algorithm",
	Short: "Returns information about a configured model algorithm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getConfiguredModelAlgorithmCmd).Standalone()

	cleanroomsml_getConfiguredModelAlgorithmCmd.Flags().String("configured-model-algorithm-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm that you want to return information about.")
	cleanroomsml_getConfiguredModelAlgorithmCmd.MarkFlagRequired("configured-model-algorithm-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getConfiguredModelAlgorithmCmd)
}
