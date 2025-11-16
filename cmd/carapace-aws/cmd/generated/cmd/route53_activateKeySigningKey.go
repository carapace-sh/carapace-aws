package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_activateKeySigningKeyCmd = &cobra.Command{
	Use:   "activate-key-signing-key",
	Short: "Activates a key-signing key (KSK) so that it can be used for signing by DNSSEC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_activateKeySigningKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_activateKeySigningKeyCmd).Standalone()

		route53_activateKeySigningKeyCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
		route53_activateKeySigningKeyCmd.Flags().String("name", "", "A string used to identify a key-signing key (KSK).")
		route53_activateKeySigningKeyCmd.MarkFlagRequired("hosted-zone-id")
		route53_activateKeySigningKeyCmd.MarkFlagRequired("name")
	})
	route53Cmd.AddCommand(route53_activateKeySigningKeyCmd)
}
