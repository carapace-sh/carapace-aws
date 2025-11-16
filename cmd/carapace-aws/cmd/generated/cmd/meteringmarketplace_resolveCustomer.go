package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var meteringmarketplace_resolveCustomerCmd = &cobra.Command{
	Use:   "resolve-customer",
	Short: "`ResolveCustomer` is called by a SaaS application during the registration process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(meteringmarketplace_resolveCustomerCmd).Standalone()

	meteringmarketplace_resolveCustomerCmd.Flags().String("registration-token", "", "When a buyer visits your website during the registration process, the buyer submits a registration token through the browser.")
	meteringmarketplace_resolveCustomerCmd.MarkFlagRequired("registration-token")
	meteringmarketplaceCmd.AddCommand(meteringmarketplace_resolveCustomerCmd)
}
