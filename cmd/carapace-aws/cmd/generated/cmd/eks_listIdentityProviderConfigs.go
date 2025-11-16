package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listIdentityProviderConfigsCmd = &cobra.Command{
	Use:   "list-identity-provider-configs",
	Short: "Lists the identity provider configurations for your cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listIdentityProviderConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listIdentityProviderConfigsCmd).Standalone()

		eks_listIdentityProviderConfigsCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_listIdentityProviderConfigsCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_listIdentityProviderConfigsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
		eks_listIdentityProviderConfigsCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_listIdentityProviderConfigsCmd)
}
