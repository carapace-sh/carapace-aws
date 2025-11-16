package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_decodeAuthorizationMessageCmd = &cobra.Command{
	Use:   "decode-authorization-message",
	Short: "Decodes additional information about the authorization status of a request from an encoded message returned in response to an Amazon Web Services request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_decodeAuthorizationMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_decodeAuthorizationMessageCmd).Standalone()

		sts_decodeAuthorizationMessageCmd.Flags().String("encoded-message", "", "The encoded message that was returned with the response.")
		sts_decodeAuthorizationMessageCmd.MarkFlagRequired("encoded-message")
	})
	stsCmd.AddCommand(sts_decodeAuthorizationMessageCmd)
}
