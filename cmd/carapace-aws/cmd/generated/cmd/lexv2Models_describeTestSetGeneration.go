package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeTestSetGenerationCmd = &cobra.Command{
	Use:   "describe-test-set-generation",
	Short: "Gets metadata information about the test set generation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeTestSetGenerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeTestSetGenerationCmd).Standalone()

		lexv2Models_describeTestSetGenerationCmd.Flags().String("test-set-generation-id", "", "The unique identifier of the test set generation.")
		lexv2Models_describeTestSetGenerationCmd.MarkFlagRequired("test-set-generation-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeTestSetGenerationCmd)
}
