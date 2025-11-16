package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getDomainAssociationCmd = &cobra.Command{
	Use:   "get-domain-association",
	Short: "Returns the domain information for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getDomainAssociationCmd).Standalone()

	amplify_getDomainAssociationCmd.Flags().String("app-id", "", "The unique id for an Amplify app.")
	amplify_getDomainAssociationCmd.Flags().String("domain-name", "", "The name of the domain.")
	amplify_getDomainAssociationCmd.MarkFlagRequired("app-id")
	amplify_getDomainAssociationCmd.MarkFlagRequired("domain-name")
	amplifyCmd.AddCommand(amplify_getDomainAssociationCmd)
}
