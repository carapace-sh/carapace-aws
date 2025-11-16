package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createKeySigningKeyCmd = &cobra.Command{
	Use:   "create-key-signing-key",
	Short: "Creates a new key-signing key (KSK) associated with a hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createKeySigningKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_createKeySigningKeyCmd).Standalone()

		route53_createKeySigningKeyCmd.Flags().String("caller-reference", "", "A unique string that identifies the request.")
		route53_createKeySigningKeyCmd.Flags().String("hosted-zone-id", "", "The unique string (ID) used to identify a hosted zone.")
		route53_createKeySigningKeyCmd.Flags().String("key-management-service-arn", "", "The Amazon resource name (ARN) for a customer managed key in Key Management Service (KMS).")
		route53_createKeySigningKeyCmd.Flags().String("name", "", "A string used to identify a key-signing key (KSK).")
		route53_createKeySigningKeyCmd.Flags().String("status", "", "A string specifying the initial status of the key-signing key (KSK).")
		route53_createKeySigningKeyCmd.MarkFlagRequired("caller-reference")
		route53_createKeySigningKeyCmd.MarkFlagRequired("hosted-zone-id")
		route53_createKeySigningKeyCmd.MarkFlagRequired("key-management-service-arn")
		route53_createKeySigningKeyCmd.MarkFlagRequired("name")
		route53_createKeySigningKeyCmd.MarkFlagRequired("status")
	})
	route53Cmd.AddCommand(route53_createKeySigningKeyCmd)
}
