package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_signPayloadCmd = &cobra.Command{
	Use:   "sign-payload",
	Short: "Signs a binary payload and returns a signature envelope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_signPayloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_signPayloadCmd).Standalone()

		signer_signPayloadCmd.Flags().String("payload", "", "Specifies the object digest (hash) to sign.")
		signer_signPayloadCmd.Flags().String("payload-format", "", "Payload content type.")
		signer_signPayloadCmd.Flags().String("profile-name", "", "The name of the signing profile.")
		signer_signPayloadCmd.Flags().String("profile-owner", "", "The AWS account ID of the profile owner.")
		signer_signPayloadCmd.MarkFlagRequired("payload")
		signer_signPayloadCmd.MarkFlagRequired("payload-format")
		signer_signPayloadCmd.MarkFlagRequired("profile-name")
	})
	signerCmd.AddCommand(signer_signPayloadCmd)
}
