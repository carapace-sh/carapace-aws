package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteOrganizationCmd = &cobra.Command{
	Use:   "delete-organization",
	Short: "Deletes an WorkMail organization and all underlying AWS resources managed by WorkMail as part of the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteOrganizationCmd).Standalone()

	workmail_deleteOrganizationCmd.Flags().String("client-token", "", "The idempotency token associated with the request.")
	workmail_deleteOrganizationCmd.Flags().Bool("delete-directory", false, "If true, deletes the AWS Directory Service directory associated with the organization.")
	workmail_deleteOrganizationCmd.Flags().Bool("delete-identity-center-application", false, "Deletes IAM Identity Center application for WorkMail.")
	workmail_deleteOrganizationCmd.Flags().Bool("force-delete", false, "Deletes a WorkMail organization even if the organization has enabled users.")
	workmail_deleteOrganizationCmd.Flags().Bool("no-delete-directory", false, "If true, deletes the AWS Directory Service directory associated with the organization.")
	workmail_deleteOrganizationCmd.Flags().Bool("no-delete-identity-center-application", false, "Deletes IAM Identity Center application for WorkMail.")
	workmail_deleteOrganizationCmd.Flags().Bool("no-force-delete", false, "Deletes a WorkMail organization even if the organization has enabled users.")
	workmail_deleteOrganizationCmd.Flags().String("organization-id", "", "The organization ID.")
	workmail_deleteOrganizationCmd.MarkFlagRequired("delete-directory")
	workmail_deleteOrganizationCmd.Flag("no-delete-directory").Hidden = true
	workmail_deleteOrganizationCmd.MarkFlagRequired("no-delete-directory")
	workmail_deleteOrganizationCmd.Flag("no-delete-identity-center-application").Hidden = true
	workmail_deleteOrganizationCmd.Flag("no-force-delete").Hidden = true
	workmail_deleteOrganizationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteOrganizationCmd)
}
