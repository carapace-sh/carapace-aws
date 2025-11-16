package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deleteChangeSetCmd = &cobra.Command{
	Use:   "delete-change-set",
	Short: "Deletes the specified change set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deleteChangeSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_deleteChangeSetCmd).Standalone()

		cloudformation_deleteChangeSetCmd.Flags().String("change-set-name", "", "The name or Amazon Resource Name (ARN) of the change set that you want to delete.")
		cloudformation_deleteChangeSetCmd.Flags().String("stack-name", "", "If you specified the name of a change set to delete, specify the stack name or Amazon Resource Name (ARN) that's associated with it.")
		cloudformation_deleteChangeSetCmd.MarkFlagRequired("change-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_deleteChangeSetCmd)
}
