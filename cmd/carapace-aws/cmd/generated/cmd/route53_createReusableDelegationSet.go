package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createReusableDelegationSetCmd = &cobra.Command{
	Use:   "create-reusable-delegation-set",
	Short: "Creates a delegation set (a group of four name servers) that can be reused by multiple hosted zones that were created by the same Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createReusableDelegationSetCmd).Standalone()

	route53_createReusableDelegationSetCmd.Flags().String("caller-reference", "", "A unique string that identifies the request, and that allows you to retry failed `CreateReusableDelegationSet` requests without the risk of executing the operation twice.")
	route53_createReusableDelegationSetCmd.Flags().String("hosted-zone-id", "", "If you want to mark the delegation set for an existing hosted zone as reusable, the ID for that hosted zone.")
	route53_createReusableDelegationSetCmd.MarkFlagRequired("caller-reference")
	route53Cmd.AddCommand(route53_createReusableDelegationSetCmd)
}
