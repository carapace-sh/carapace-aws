package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteKeySigningKeyCmd = &cobra.Command{
	Use:   "delete-key-signing-key",
	Short: "Deletes a key-signing key (KSK).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteKeySigningKeyCmd).Standalone()

	route53_deleteKeySigningKeyCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
	route53_deleteKeySigningKeyCmd.Flags().String("name", "", "A string used to identify a key-signing key (KSK).")
	route53_deleteKeySigningKeyCmd.MarkFlagRequired("hosted-zone-id")
	route53_deleteKeySigningKeyCmd.MarkFlagRequired("name")
	route53Cmd.AddCommand(route53_deleteKeySigningKeyCmd)
}
