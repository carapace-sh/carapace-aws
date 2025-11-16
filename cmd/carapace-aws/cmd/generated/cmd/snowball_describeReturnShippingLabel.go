package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_describeReturnShippingLabelCmd = &cobra.Command{
	Use:   "describe-return-shipping-label",
	Short: "Information on the shipping label of a Snow device that is being returned to Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_describeReturnShippingLabelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_describeReturnShippingLabelCmd).Standalone()

		snowball_describeReturnShippingLabelCmd.Flags().String("job-id", "", "The automatically generated ID for a job, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_describeReturnShippingLabelCmd.MarkFlagRequired("job-id")
	})
	snowballCmd.AddCommand(snowball_describeReturnShippingLabelCmd)
}
