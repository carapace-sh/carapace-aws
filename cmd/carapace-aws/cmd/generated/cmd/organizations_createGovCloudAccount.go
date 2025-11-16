package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_createGovCloudAccountCmd = &cobra.Command{
	Use:   "create-gov-cloud-account",
	Short: "This action is available if all of the following are true:\n\n- You're authorized to create accounts in the Amazon Web Services GovCloud (US) Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createGovCloudAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_createGovCloudAccountCmd).Standalone()

		organizations_createGovCloudAccountCmd.Flags().String("account-name", "", "The friendly name of the member account.")
		organizations_createGovCloudAccountCmd.Flags().String("email", "", "Specifies the email address of the owner to assign to the new member account in the commercial Region.")
		organizations_createGovCloudAccountCmd.Flags().String("iam-user-access-to-billing", "", "If set to `ALLOW`, the new linked account in the commercial Region enables IAM users to access account billing information *if* they have the required permissions.")
		organizations_createGovCloudAccountCmd.Flags().String("role-name", "", "(Optional)")
		organizations_createGovCloudAccountCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created account.")
		organizations_createGovCloudAccountCmd.MarkFlagRequired("account-name")
		organizations_createGovCloudAccountCmd.MarkFlagRequired("email")
	})
	organizationsCmd.AddCommand(organizations_createGovCloudAccountCmd)
}
