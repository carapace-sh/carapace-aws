package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_updateInvestigationGroupCmd = &cobra.Command{
	Use:   "update-investigation-group",
	Short: "Updates the configuration of the specified investigation group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_updateInvestigationGroupCmd).Standalone()

	aiops_updateInvestigationGroupCmd.Flags().String("chatbot-notification-channel", "", "Use this structure to integrate CloudWatch investigations with chat applications.")
	aiops_updateInvestigationGroupCmd.Flags().String("cross-account-configurations", "", "Used to configure cross-account access for an investigation group.")
	aiops_updateInvestigationGroupCmd.Flags().String("encryption-configuration", "", "Use this structure if you want to use a customer managed KMS key to encrypt your investigation data.")
	aiops_updateInvestigationGroupCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to modify.")
	aiops_updateInvestigationGroupCmd.Flags().Bool("is-cloud-trail-event-history-enabled", false, "Specify `true` to enable CloudWatch investigations to have access to change events that are recorded by CloudTrail.")
	aiops_updateInvestigationGroupCmd.Flags().Bool("no-is-cloud-trail-event-history-enabled", false, "Specify `true` to enable CloudWatch investigations to have access to change events that are recorded by CloudTrail.")
	aiops_updateInvestigationGroupCmd.Flags().String("role-arn", "", "Specify this field if you want to change the IAM role that CloudWatch investigations will use when it gathers investigation data.")
	aiops_updateInvestigationGroupCmd.Flags().String("tag-key-boundaries", "", "Enter the existing custom tag keys for custom applications in your system.")
	aiops_updateInvestigationGroupCmd.MarkFlagRequired("identifier")
	aiops_updateInvestigationGroupCmd.Flag("no-is-cloud-trail-event-history-enabled").Hidden = true
	aiopsCmd.AddCommand(aiops_updateInvestigationGroupCmd)
}
