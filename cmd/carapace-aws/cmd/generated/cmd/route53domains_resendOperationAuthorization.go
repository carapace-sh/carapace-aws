package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_resendOperationAuthorizationCmd = &cobra.Command{
	Use:   "resend-operation-authorization",
	Short: "Resend the form of authorization email for this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_resendOperationAuthorizationCmd).Standalone()

	route53domains_resendOperationAuthorizationCmd.Flags().String("operation-id", "", "Operation ID.")
	route53domains_resendOperationAuthorizationCmd.MarkFlagRequired("operation-id")
	route53domainsCmd.AddCommand(route53domains_resendOperationAuthorizationCmd)
}
