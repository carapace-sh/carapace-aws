package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createMitigationActionCmd = &cobra.Command{
	Use:   "create-mitigation-action",
	Short: "Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createMitigationActionCmd).Standalone()

	iot_createMitigationActionCmd.Flags().String("action-name", "", "A friendly name for the action.")
	iot_createMitigationActionCmd.Flags().String("action-params", "", "Defines the type of action and the parameters for that action.")
	iot_createMitigationActionCmd.Flags().String("role-arn", "", "The ARN of the IAM role that is used to apply the mitigation action.")
	iot_createMitigationActionCmd.Flags().String("tags", "", "Metadata that can be used to manage the mitigation action.")
	iot_createMitigationActionCmd.MarkFlagRequired("action-name")
	iot_createMitigationActionCmd.MarkFlagRequired("action-params")
	iot_createMitigationActionCmd.MarkFlagRequired("role-arn")
	iotCmd.AddCommand(iot_createMitigationActionCmd)
}
