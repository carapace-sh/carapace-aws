package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getIdentityCenterAuthTokenCmd = &cobra.Command{
	Use:   "get-identity-center-auth-token",
	Short: "Generates an encrypted authentication token that propagates the caller's Amazon Web Services IAM Identity Center identity to Amazon Redshift clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getIdentityCenterAuthTokenCmd).Standalone()

	redshift_getIdentityCenterAuthTokenCmd.Flags().String("cluster-ids", "", "A list of cluster identifiers that the generated token can be used with.")
	redshift_getIdentityCenterAuthTokenCmd.MarkFlagRequired("cluster-ids")
	redshiftCmd.AddCommand(redshift_getIdentityCenterAuthTokenCmd)
}
