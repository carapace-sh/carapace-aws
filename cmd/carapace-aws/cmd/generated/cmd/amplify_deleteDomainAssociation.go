package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteDomainAssociationCmd = &cobra.Command{
	Use:   "delete-domain-association",
	Short: "Deletes a domain association for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteDomainAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_deleteDomainAssociationCmd).Standalone()

		amplify_deleteDomainAssociationCmd.Flags().String("app-id", "", "The unique id for an Amplify app.")
		amplify_deleteDomainAssociationCmd.Flags().String("domain-name", "", "The name of the domain.")
		amplify_deleteDomainAssociationCmd.MarkFlagRequired("app-id")
		amplify_deleteDomainAssociationCmd.MarkFlagRequired("domain-name")
	})
	amplifyCmd.AddCommand(amplify_deleteDomainAssociationCmd)
}
