package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeTestSetCmd = &cobra.Command{
	Use:   "describe-test-set",
	Short: "Gets metadata information about the test set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeTestSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeTestSetCmd).Standalone()

		lexv2Models_describeTestSetCmd.Flags().String("test-set-id", "", "The test set Id for the test set request.")
		lexv2Models_describeTestSetCmd.MarkFlagRequired("test-set-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeTestSetCmd)
}
