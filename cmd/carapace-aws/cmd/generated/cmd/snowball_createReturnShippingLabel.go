package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_createReturnShippingLabelCmd = &cobra.Command{
	Use:   "create-return-shipping-label",
	Short: "Creates a shipping label that will be used to return the Snow device to Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_createReturnShippingLabelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_createReturnShippingLabelCmd).Standalone()

		snowball_createReturnShippingLabelCmd.Flags().String("job-id", "", "The ID for a job that you want to create the return shipping label for; for example, `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_createReturnShippingLabelCmd.Flags().String("shipping-option", "", "The shipping speed for a particular job.")
		snowball_createReturnShippingLabelCmd.MarkFlagRequired("job-id")
	})
	snowballCmd.AddCommand(snowball_createReturnShippingLabelCmd)
}
