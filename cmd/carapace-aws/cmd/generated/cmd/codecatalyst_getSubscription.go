package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getSubscriptionCmd = &cobra.Command{
	Use:   "get-subscription",
	Short: "Returns information about the Amazon Web Services account used for billing purposes and the billing plan for the space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getSubscriptionCmd).Standalone()

	codecatalyst_getSubscriptionCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_getSubscriptionCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_getSubscriptionCmd)
}
