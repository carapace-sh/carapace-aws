package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_createInvestigationGroupCmd = &cobra.Command{
	Use:   "create-investigation-group",
	Short: "Creates an *investigation group* in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_createInvestigationGroupCmd).Standalone()

	aiops_createInvestigationGroupCmd.Flags().String("chatbot-notification-channel", "", "Use this structure to integrate CloudWatch investigations with chat applications.")
	aiops_createInvestigationGroupCmd.Flags().String("cross-account-configurations", "", "List of `sourceRoleArn` values that have been configured for cross-account access.")
	aiops_createInvestigationGroupCmd.Flags().String("encryption-configuration", "", "Use this structure if you want to use a customer managed KMS key to encrypt your investigation data.")
	aiops_createInvestigationGroupCmd.Flags().Bool("is-cloud-trail-event-history-enabled", false, "Specify `true` to enable CloudWatch investigations to have access to change events that are recorded by CloudTrail.")
	aiops_createInvestigationGroupCmd.Flags().String("name", "", "Provides a name for the investigation group.")
	aiops_createInvestigationGroupCmd.Flags().Bool("no-is-cloud-trail-event-history-enabled", false, "Specify `true` to enable CloudWatch investigations to have access to change events that are recorded by CloudTrail.")
	aiops_createInvestigationGroupCmd.Flags().String("retention-in-days", "", "Specify how long that investigation data is kept.")
	aiops_createInvestigationGroupCmd.Flags().String("role-arn", "", "Specify the ARN of the IAM role that CloudWatch investigations will use when it gathers investigation data.")
	aiops_createInvestigationGroupCmd.Flags().String("tag-key-boundaries", "", "Enter the existing custom tag keys for custom applications in your system.")
	aiops_createInvestigationGroupCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the investigation group.")
	aiops_createInvestigationGroupCmd.MarkFlagRequired("name")
	aiops_createInvestigationGroupCmd.Flag("no-is-cloud-trail-event-history-enabled").Hidden = true
	aiops_createInvestigationGroupCmd.MarkFlagRequired("role-arn")
	aiopsCmd.AddCommand(aiops_createInvestigationGroupCmd)
}
