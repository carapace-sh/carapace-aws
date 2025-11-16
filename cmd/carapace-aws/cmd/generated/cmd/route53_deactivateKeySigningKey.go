package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deactivateKeySigningKeyCmd = &cobra.Command{
	Use:   "deactivate-key-signing-key",
	Short: "Deactivates a key-signing key (KSK) so that it will not be used for signing by DNSSEC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deactivateKeySigningKeyCmd).Standalone()

	route53_deactivateKeySigningKeyCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
	route53_deactivateKeySigningKeyCmd.Flags().String("name", "", "A string used to identify a key-signing key (KSK).")
	route53_deactivateKeySigningKeyCmd.MarkFlagRequired("hosted-zone-id")
	route53_deactivateKeySigningKeyCmd.MarkFlagRequired("name")
	route53Cmd.AddCommand(route53_deactivateKeySigningKeyCmd)
}
