package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_putSigningProfileCmd = &cobra.Command{
	Use:   "put-signing-profile",
	Short: "Creates a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_putSigningProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_putSigningProfileCmd).Standalone()

		signer_putSigningProfileCmd.Flags().String("overrides", "", "A subfield of `platform`.")
		signer_putSigningProfileCmd.Flags().String("platform-id", "", "The ID of the signing platform to be created.")
		signer_putSigningProfileCmd.Flags().String("profile-name", "", "The name of the signing profile to be created.")
		signer_putSigningProfileCmd.Flags().String("signature-validity-period", "", "The default validity period override for any signature generated using this signing profile.")
		signer_putSigningProfileCmd.Flags().String("signing-material", "", "The AWS Certificate Manager certificate that will be used to sign code with the new signing profile.")
		signer_putSigningProfileCmd.Flags().String("signing-parameters", "", "Map of key-value pairs for signing.")
		signer_putSigningProfileCmd.Flags().String("tags", "", "Tags to be associated with the signing profile that is being created.")
		signer_putSigningProfileCmd.MarkFlagRequired("platform-id")
		signer_putSigningProfileCmd.MarkFlagRequired("profile-name")
	})
	signerCmd.AddCommand(signer_putSigningProfileCmd)
}
