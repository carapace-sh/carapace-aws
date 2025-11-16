package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putInboundDmarcSettingsCmd = &cobra.Command{
	Use:   "put-inbound-dmarc-settings",
	Short: "Enables or disables a DMARC policy for a given organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putInboundDmarcSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_putInboundDmarcSettingsCmd).Standalone()

		workmail_putInboundDmarcSettingsCmd.Flags().String("enforced", "", "Enforces or suspends a policy after it's applied.")
		workmail_putInboundDmarcSettingsCmd.Flags().String("organization-id", "", "The ID of the organization that you are applying the DMARC policy to.")
		workmail_putInboundDmarcSettingsCmd.MarkFlagRequired("enforced")
		workmail_putInboundDmarcSettingsCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_putInboundDmarcSettingsCmd)
}
