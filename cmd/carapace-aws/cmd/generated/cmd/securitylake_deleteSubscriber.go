package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteSubscriberCmd = &cobra.Command{
	Use:   "delete-subscriber",
	Short: "Deletes the subscription permission and all notification settings for accounts that are already enabled in Amazon Security Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteSubscriberCmd).Standalone()

	securitylake_deleteSubscriberCmd.Flags().String("subscriber-id", "", "A value created by Security Lake that uniquely identifies your `DeleteSubscriber` API request.")
	securitylake_deleteSubscriberCmd.MarkFlagRequired("subscriber-id")
	securitylakeCmd.AddCommand(securitylake_deleteSubscriberCmd)
}
