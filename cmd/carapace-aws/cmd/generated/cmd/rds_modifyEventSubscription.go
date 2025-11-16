package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyEventSubscriptionCmd = &cobra.Command{
	Use:   "modify-event-subscription",
	Short: "Modifies an existing RDS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyEventSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyEventSubscriptionCmd).Standalone()

		rds_modifyEventSubscriptionCmd.Flags().String("enabled", "", "Specifies whether to activate the subscription.")
		rds_modifyEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a source type (`SourceType`) that you want to subscribe to.")
		rds_modifyEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
		rds_modifyEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
		rds_modifyEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the RDS event notification subscription.")
		rds_modifyEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	rdsCmd.AddCommand(rds_modifyEventSubscriptionCmd)
}
