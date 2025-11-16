package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_registerOrganizationDelegatedAdminCmd = &cobra.Command{
	Use:   "register-organization-delegated-admin",
	Short: "Registers an organization’s member account as the CloudTrail [delegated administrator](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-delegated-administrator.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_registerOrganizationDelegatedAdminCmd).Standalone()

	cloudtrail_registerOrganizationDelegatedAdminCmd.Flags().String("member-account-id", "", "An organization member account ID that you want to designate as a delegated administrator.")
	cloudtrail_registerOrganizationDelegatedAdminCmd.MarkFlagRequired("member-account-id")
	cloudtrailCmd.AddCommand(cloudtrail_registerOrganizationDelegatedAdminCmd)
}
