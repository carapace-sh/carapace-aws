package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a domain that contains all Amazon Connect Voice ID data, such as speakers, fraudsters, customer audio, and voiceprints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_createDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_createDomainCmd).Standalone()

		voiceId_createDomainCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		voiceId_createDomainCmd.Flags().String("description", "", "A brief description of this domain.")
		voiceId_createDomainCmd.Flags().String("name", "", "The name of the domain.")
		voiceId_createDomainCmd.Flags().String("server-side-encryption-configuration", "", "The configuration, containing the KMS key identifier, to be used by Voice ID for the server-side encryption of your data.")
		voiceId_createDomainCmd.Flags().String("tags", "", "A list of tags you want added to the domain.")
		voiceId_createDomainCmd.MarkFlagRequired("name")
		voiceId_createDomainCmd.MarkFlagRequired("server-side-encryption-configuration")
	})
	voiceIdCmd.AddCommand(voiceId_createDomainCmd)
}
