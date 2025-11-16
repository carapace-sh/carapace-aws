package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_executeChangeSetCmd = &cobra.Command{
	Use:   "execute-change-set",
	Short: "Updates a stack using the input information that was provided when the specified change set was created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_executeChangeSetCmd).Standalone()

	cloudformation_executeChangeSetCmd.Flags().String("change-set-name", "", "The name or Amazon Resource Name (ARN) of the change set that you want use to update the specified stack.")
	cloudformation_executeChangeSetCmd.Flags().String("client-request-token", "", "A unique identifier for this `ExecuteChangeSet` request.")
	cloudformation_executeChangeSetCmd.Flags().String("disable-rollback", "", "Preserves the state of previously provisioned resources when an operation fails.")
	cloudformation_executeChangeSetCmd.Flags().String("retain-except-on-create", "", "When set to `true`, newly created resources are deleted when the operation rolls back.")
	cloudformation_executeChangeSetCmd.Flags().String("stack-name", "", "If you specified the name of a change set, specify the stack name or Amazon Resource Name (ARN) that's associated with the change set you want to execute.")
	cloudformation_executeChangeSetCmd.MarkFlagRequired("change-set-name")
	cloudformationCmd.AddCommand(cloudformation_executeChangeSetCmd)
}
