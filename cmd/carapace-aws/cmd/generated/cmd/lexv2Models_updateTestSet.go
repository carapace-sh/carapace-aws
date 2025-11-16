package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateTestSetCmd = &cobra.Command{
	Use:   "update-test-set",
	Short: "The action to update the test set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateTestSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateTestSetCmd).Standalone()

		lexv2Models_updateTestSetCmd.Flags().String("description", "", "The new test set description.")
		lexv2Models_updateTestSetCmd.Flags().String("test-set-id", "", "The test set Id for which update test operation to be performed.")
		lexv2Models_updateTestSetCmd.Flags().String("test-set-name", "", "The new test set name.")
		lexv2Models_updateTestSetCmd.MarkFlagRequired("test-set-id")
		lexv2Models_updateTestSetCmd.MarkFlagRequired("test-set-name")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateTestSetCmd)
}
