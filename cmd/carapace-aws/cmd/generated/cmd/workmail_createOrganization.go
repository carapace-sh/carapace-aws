package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createOrganizationCmd = &cobra.Command{
	Use:   "create-organization",
	Short: "Creates a new WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createOrganizationCmd).Standalone()

	workmail_createOrganizationCmd.Flags().String("alias", "", "The organization alias.")
	workmail_createOrganizationCmd.Flags().String("client-token", "", "The idempotency token associated with the request.")
	workmail_createOrganizationCmd.Flags().String("directory-id", "", "The AWS Directory Service directory ID.")
	workmail_createOrganizationCmd.Flags().String("domains", "", "The email domains to associate with the organization.")
	workmail_createOrganizationCmd.Flags().Bool("enable-interoperability", false, "When `true`, allows organization interoperability between WorkMail and Microsoft Exchange.")
	workmail_createOrganizationCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of a customer managed key from AWS KMS.")
	workmail_createOrganizationCmd.Flags().Bool("no-enable-interoperability", false, "When `true`, allows organization interoperability between WorkMail and Microsoft Exchange.")
	workmail_createOrganizationCmd.MarkFlagRequired("alias")
	workmail_createOrganizationCmd.Flag("no-enable-interoperability").Hidden = true
	workmailCmd.AddCommand(workmail_createOrganizationCmd)
}
