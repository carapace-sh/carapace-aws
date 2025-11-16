package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteTestSetCmd = &cobra.Command{
	Use:   "delete-test-set",
	Short: "The action to delete the selected test set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteTestSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteTestSetCmd).Standalone()

		lexv2Models_deleteTestSetCmd.Flags().String("test-set-id", "", "The test set Id of the test set to be deleted.")
		lexv2Models_deleteTestSetCmd.MarkFlagRequired("test-set-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteTestSetCmd)
}
