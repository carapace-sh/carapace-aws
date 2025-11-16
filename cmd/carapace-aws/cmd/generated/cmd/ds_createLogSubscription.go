package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createLogSubscriptionCmd = &cobra.Command{
	Use:   "create-log-subscription",
	Short: "Creates a subscription to forward real-time Directory Service domain controller security logs to the specified Amazon CloudWatch log group in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createLogSubscriptionCmd).Standalone()

	ds_createLogSubscriptionCmd.Flags().String("directory-id", "", "Identifier of the directory to which you want to subscribe and receive real-time logs to your specified CloudWatch log group.")
	ds_createLogSubscriptionCmd.Flags().String("log-group-name", "", "The name of the CloudWatch log group where the real-time domain controller logs are forwarded.")
	ds_createLogSubscriptionCmd.MarkFlagRequired("directory-id")
	ds_createLogSubscriptionCmd.MarkFlagRequired("log-group-name")
	dsCmd.AddCommand(ds_createLogSubscriptionCmd)
}
