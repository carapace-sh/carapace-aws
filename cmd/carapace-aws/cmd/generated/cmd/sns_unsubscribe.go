package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "Deletes a subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_unsubscribeCmd).Standalone()

	sns_unsubscribeCmd.Flags().String("subscription-arn", "", "The ARN of the subscription to be deleted.")
	sns_unsubscribeCmd.MarkFlagRequired("subscription-arn")
	snsCmd.AddCommand(sns_unsubscribeCmd)
}
