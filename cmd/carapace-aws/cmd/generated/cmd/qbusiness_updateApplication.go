package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates an existing Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_updateApplicationCmd).Standalone()

		qbusiness_updateApplicationCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
		qbusiness_updateApplicationCmd.Flags().String("attachments-configuration", "", "An option to allow end users to upload files directly during chat.")
		qbusiness_updateApplicationCmd.Flags().String("auto-subscription-configuration", "", "An option to enable updating the default subscription type assigned to an Amazon Q Business application using IAM identity federation for user management.")
		qbusiness_updateApplicationCmd.Flags().String("description", "", "A description for the Amazon Q Business application.")
		qbusiness_updateApplicationCmd.Flags().String("display-name", "", "A name for the Amazon Q Business application.")
		qbusiness_updateApplicationCmd.Flags().String("identity-center-instance-arn", "", "The Amazon Resource Name (ARN) of the IAM Identity Center instance you are either creating for—or connecting to—your Amazon Q Business application.")
		qbusiness_updateApplicationCmd.Flags().String("personalization-configuration", "", "Configuration information about chat response personalization.")
		qbusiness_updateApplicationCmd.Flags().String("q-apps-configuration", "", "An option to allow end users to create and use Amazon Q Apps in the web experience.")
		qbusiness_updateApplicationCmd.Flags().String("role-arn", "", "An Amazon Web Services Identity and Access Management (IAM) role that gives Amazon Q Business permission to access Amazon CloudWatch logs and metrics.")
		qbusiness_updateApplicationCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_updateApplicationCmd)
}
