package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateDefaultMailDomainCmd = &cobra.Command{
	Use:   "update-default-mail-domain",
	Short: "Updates the default mail domain for an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateDefaultMailDomainCmd).Standalone()

	workmail_updateDefaultMailDomainCmd.Flags().String("domain-name", "", "The domain name that will become the default domain.")
	workmail_updateDefaultMailDomainCmd.Flags().String("organization-id", "", "The WorkMail organization for which to list domains.")
	workmail_updateDefaultMailDomainCmd.MarkFlagRequired("domain-name")
	workmail_updateDefaultMailDomainCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_updateDefaultMailDomainCmd)
}
