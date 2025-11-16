package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_updateSubscriberCmd = &cobra.Command{
	Use:   "update-subscriber",
	Short: "Updates an existing subscription for the given Amazon Security Lake account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_updateSubscriberCmd).Standalone()

	securitylake_updateSubscriberCmd.Flags().String("sources", "", "The supported Amazon Web Services services from which logs and events are collected.")
	securitylake_updateSubscriberCmd.Flags().String("subscriber-description", "", "The description of the Security Lake account subscriber.")
	securitylake_updateSubscriberCmd.Flags().String("subscriber-id", "", "A value created by Security Lake that uniquely identifies your subscription.")
	securitylake_updateSubscriberCmd.Flags().String("subscriber-identity", "", "The Amazon Web Services identity used to access your data.")
	securitylake_updateSubscriberCmd.Flags().String("subscriber-name", "", "The name of the Security Lake account subscriber.")
	securitylake_updateSubscriberCmd.MarkFlagRequired("subscriber-id")
	securitylakeCmd.AddCommand(securitylake_updateSubscriberCmd)
}
