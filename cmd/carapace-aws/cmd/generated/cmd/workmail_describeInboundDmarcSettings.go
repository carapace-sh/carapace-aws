package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeInboundDmarcSettingsCmd = &cobra.Command{
	Use:   "describe-inbound-dmarc-settings",
	Short: "Lists the settings in a DMARC policy for a specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeInboundDmarcSettingsCmd).Standalone()

	workmail_describeInboundDmarcSettingsCmd.Flags().String("organization-id", "", "Lists the ID of the given organization.")
	workmail_describeInboundDmarcSettingsCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_describeInboundDmarcSettingsCmd)
}
