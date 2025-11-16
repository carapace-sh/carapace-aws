package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_getSubscriberCmd = &cobra.Command{
	Use:   "get-subscriber",
	Short: "Retrieves the subscription information for the specified subscription ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_getSubscriberCmd).Standalone()

	securitylake_getSubscriberCmd.Flags().String("subscriber-id", "", "A value created by Amazon Security Lake that uniquely identifies your `GetSubscriber` API request.")
	securitylake_getSubscriberCmd.MarkFlagRequired("subscriber-id")
	securitylakeCmd.AddCommand(securitylake_getSubscriberCmd)
}
