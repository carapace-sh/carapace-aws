package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getApiAssociationCmd = &cobra.Command{
	Use:   "get-api-association",
	Short: "Retrieves an `ApiAssociation` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getApiAssociationCmd).Standalone()

	appsync_getApiAssociationCmd.Flags().String("domain-name", "", "The domain name.")
	appsync_getApiAssociationCmd.MarkFlagRequired("domain-name")
	appsyncCmd.AddCommand(appsync_getApiAssociationCmd)
}
