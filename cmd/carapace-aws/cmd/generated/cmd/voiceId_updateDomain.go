package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_updateDomainCmd = &cobra.Command{
	Use:   "update-domain",
	Short: "Updates the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_updateDomainCmd).Standalone()

	voiceId_updateDomainCmd.Flags().String("description", "", "A brief description about this domain.")
	voiceId_updateDomainCmd.Flags().String("domain-id", "", "The identifier of the domain to be updated.")
	voiceId_updateDomainCmd.Flags().String("name", "", "The name of the domain.")
	voiceId_updateDomainCmd.Flags().String("server-side-encryption-configuration", "", "The configuration, containing the KMS key identifier, to be used by Voice ID for the server-side encryption of your data.")
	voiceId_updateDomainCmd.MarkFlagRequired("domain-id")
	voiceId_updateDomainCmd.MarkFlagRequired("name")
	voiceId_updateDomainCmd.MarkFlagRequired("server-side-encryption-configuration")
	voiceIdCmd.AddCommand(voiceId_updateDomainCmd)
}
