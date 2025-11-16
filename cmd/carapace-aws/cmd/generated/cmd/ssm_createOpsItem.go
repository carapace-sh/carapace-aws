package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createOpsItemCmd = &cobra.Command{
	Use:   "create-ops-item",
	Short: "Creates a new OpsItem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createOpsItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_createOpsItemCmd).Standalone()

		ssm_createOpsItemCmd.Flags().String("account-id", "", "The target Amazon Web Services account where you want to create an OpsItem.")
		ssm_createOpsItemCmd.Flags().String("actual-end-time", "", "The time a runbook workflow ended.")
		ssm_createOpsItemCmd.Flags().String("actual-start-time", "", "The time a runbook workflow started.")
		ssm_createOpsItemCmd.Flags().String("category", "", "Specify a category to assign to an OpsItem.")
		ssm_createOpsItemCmd.Flags().String("description", "", "User-defined text that contains information about the OpsItem, in Markdown format.")
		ssm_createOpsItemCmd.Flags().String("notifications", "", "The Amazon Resource Name (ARN) of an SNS topic where notifications are sent when this OpsItem is edited or changed.")
		ssm_createOpsItemCmd.Flags().String("operational-data", "", "Operational data is custom data that provides useful reference details about the OpsItem.")
		ssm_createOpsItemCmd.Flags().String("ops-item-type", "", "The type of OpsItem to create.")
		ssm_createOpsItemCmd.Flags().String("planned-end-time", "", "The time specified in a change request for a runbook workflow to end.")
		ssm_createOpsItemCmd.Flags().String("planned-start-time", "", "The time specified in a change request for a runbook workflow to start.")
		ssm_createOpsItemCmd.Flags().String("priority", "", "The importance of this OpsItem in relation to other OpsItems in the system.")
		ssm_createOpsItemCmd.Flags().String("related-ops-items", "", "One or more OpsItems that share something in common with the current OpsItems.")
		ssm_createOpsItemCmd.Flags().String("severity", "", "Specify a severity to assign to an OpsItem.")
		ssm_createOpsItemCmd.Flags().String("source", "", "The origin of the OpsItem, such as Amazon EC2 or Systems Manager.")
		ssm_createOpsItemCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_createOpsItemCmd.Flags().String("title", "", "A short heading that describes the nature of the OpsItem and the impacted resource.")
		ssm_createOpsItemCmd.MarkFlagRequired("description")
		ssm_createOpsItemCmd.MarkFlagRequired("source")
		ssm_createOpsItemCmd.MarkFlagRequired("title")
	})
	ssmCmd.AddCommand(ssm_createOpsItemCmd)
}
