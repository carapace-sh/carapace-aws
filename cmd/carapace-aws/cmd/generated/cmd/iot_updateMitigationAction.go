package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateMitigationActionCmd = &cobra.Command{
	Use:   "update-mitigation-action",
	Short: "Updates the definition for the specified mitigation action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateMitigationActionCmd).Standalone()

	iot_updateMitigationActionCmd.Flags().String("action-name", "", "The friendly name for the mitigation action.")
	iot_updateMitigationActionCmd.Flags().String("action-params", "", "Defines the type of action and the parameters for that action.")
	iot_updateMitigationActionCmd.Flags().String("role-arn", "", "The ARN of the IAM role that is used to apply the mitigation action.")
	iot_updateMitigationActionCmd.MarkFlagRequired("action-name")
	iotCmd.AddCommand(iot_updateMitigationActionCmd)
}
