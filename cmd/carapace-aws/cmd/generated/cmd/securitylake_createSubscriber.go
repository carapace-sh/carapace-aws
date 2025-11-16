package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createSubscriberCmd = &cobra.Command{
	Use:   "create-subscriber",
	Short: "Creates a subscriber for accounts that are already enabled in Amazon Security Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createSubscriberCmd).Standalone()

	securitylake_createSubscriberCmd.Flags().String("access-types", "", "The Amazon S3 or Lake Formation access type.")
	securitylake_createSubscriberCmd.Flags().String("sources", "", "The supported Amazon Web Services services from which logs and events are collected.")
	securitylake_createSubscriberCmd.Flags().String("subscriber-description", "", "The description for your subscriber account in Security Lake.")
	securitylake_createSubscriberCmd.Flags().String("subscriber-identity", "", "The Amazon Web Services identity used to access your data.")
	securitylake_createSubscriberCmd.Flags().String("subscriber-name", "", "The name of your Security Lake subscriber account.")
	securitylake_createSubscriberCmd.Flags().String("tags", "", "An array of objects, one for each tag to associate with the subscriber.")
	securitylake_createSubscriberCmd.MarkFlagRequired("sources")
	securitylake_createSubscriberCmd.MarkFlagRequired("subscriber-identity")
	securitylake_createSubscriberCmd.MarkFlagRequired("subscriber-name")
	securitylakeCmd.AddCommand(securitylake_createSubscriberCmd)
}
