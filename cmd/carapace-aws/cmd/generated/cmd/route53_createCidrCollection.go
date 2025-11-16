package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createCidrCollectionCmd = &cobra.Command{
	Use:   "create-cidr-collection",
	Short: "Creates a CIDR collection in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createCidrCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_createCidrCollectionCmd).Standalone()

		route53_createCidrCollectionCmd.Flags().String("caller-reference", "", "A client-specific token that allows requests to be securely retried so that the intended outcome will only occur once, retries receive a similar response, and there are no additional edge cases to handle.")
		route53_createCidrCollectionCmd.Flags().String("name", "", "A unique identifier for the account that can be used to reference the collection from other API calls.")
		route53_createCidrCollectionCmd.MarkFlagRequired("caller-reference")
		route53_createCidrCollectionCmd.MarkFlagRequired("name")
	})
	route53Cmd.AddCommand(route53_createCidrCollectionCmd)
}
