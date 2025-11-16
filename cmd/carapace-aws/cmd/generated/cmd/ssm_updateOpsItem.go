package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateOpsItemCmd = &cobra.Command{
	Use:   "update-ops-item",
	Short: "Edit or change an OpsItem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateOpsItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateOpsItemCmd).Standalone()

		ssm_updateOpsItemCmd.Flags().String("actual-end-time", "", "The time a runbook workflow ended.")
		ssm_updateOpsItemCmd.Flags().String("actual-start-time", "", "The time a runbook workflow started.")
		ssm_updateOpsItemCmd.Flags().String("category", "", "Specify a new category for an OpsItem.")
		ssm_updateOpsItemCmd.Flags().String("description", "", "User-defined text that contains information about the OpsItem, in Markdown format.")
		ssm_updateOpsItemCmd.Flags().String("notifications", "", "The Amazon Resource Name (ARN) of an SNS topic where notifications are sent when this OpsItem is edited or changed.")
		ssm_updateOpsItemCmd.Flags().String("operational-data", "", "Add new keys or edit existing key-value pairs of the OperationalData map in the OpsItem object.")
		ssm_updateOpsItemCmd.Flags().String("operational-data-to-delete", "", "Keys that you want to remove from the OperationalData map.")
		ssm_updateOpsItemCmd.Flags().String("ops-item-arn", "", "The OpsItem Amazon Resource Name (ARN).")
		ssm_updateOpsItemCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem.")
		ssm_updateOpsItemCmd.Flags().String("planned-end-time", "", "The time specified in a change request for a runbook workflow to end.")
		ssm_updateOpsItemCmd.Flags().String("planned-start-time", "", "The time specified in a change request for a runbook workflow to start.")
		ssm_updateOpsItemCmd.Flags().String("priority", "", "The importance of this OpsItem in relation to other OpsItems in the system.")
		ssm_updateOpsItemCmd.Flags().String("related-ops-items", "", "One or more OpsItems that share something in common with the current OpsItems.")
		ssm_updateOpsItemCmd.Flags().String("severity", "", "Specify a new severity for an OpsItem.")
		ssm_updateOpsItemCmd.Flags().String("status", "", "The OpsItem status.")
		ssm_updateOpsItemCmd.Flags().String("title", "", "A short heading that describes the nature of the OpsItem and the impacted resource.")
		ssm_updateOpsItemCmd.MarkFlagRequired("ops-item-id")
	})
	ssmCmd.AddCommand(ssm_updateOpsItemCmd)
}
