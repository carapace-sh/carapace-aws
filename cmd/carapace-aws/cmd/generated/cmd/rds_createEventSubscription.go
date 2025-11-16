package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createEventSubscriptionCmd = &cobra.Command{
	Use:   "create-event-subscription",
	Short: "Creates an RDS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createEventSubscriptionCmd).Standalone()

	rds_createEventSubscriptionCmd.Flags().String("enabled", "", "Specifies whether to activate the subscription.")
	rds_createEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a particular source type (`SourceType`) that you want to subscribe to.")
	rds_createEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
	rds_createEventSubscriptionCmd.Flags().String("source-ids", "", "The list of identifiers of the event sources for which events are returned.")
	rds_createEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	rds_createEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the subscription.")
	rds_createEventSubscriptionCmd.Flags().String("tags", "", "")
	rds_createEventSubscriptionCmd.MarkFlagRequired("sns-topic-arn")
	rds_createEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	rdsCmd.AddCommand(rds_createEventSubscriptionCmd)
}
