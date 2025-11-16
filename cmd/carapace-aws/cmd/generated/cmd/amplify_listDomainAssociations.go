package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listDomainAssociationsCmd = &cobra.Command{
	Use:   "list-domain-associations",
	Short: "Returns the domain associations for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listDomainAssociationsCmd).Standalone()

	amplify_listDomainAssociationsCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_listDomainAssociationsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
	amplify_listDomainAssociationsCmd.Flags().String("next-token", "", "A pagination token.")
	amplify_listDomainAssociationsCmd.MarkFlagRequired("app-id")
	amplifyCmd.AddCommand(amplify_listDomainAssociationsCmd)
}
