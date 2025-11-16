package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createEventSubscriptionCmd = &cobra.Command{
	Use:   "create-event-subscription",
	Short: "Creates an Amazon Redshift event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createEventSubscriptionCmd).Standalone()

	redshift_createEventSubscriptionCmd.Flags().String("enabled", "", "A boolean value; set to `true` to activate the subscription, and set to `false` to create the subscription but not activate it.")
	redshift_createEventSubscriptionCmd.Flags().String("event-categories", "", "Specifies the Amazon Redshift event categories to be published by the event notification subscription.")
	redshift_createEventSubscriptionCmd.Flags().String("severity", "", "Specifies the Amazon Redshift event severity to be published by the event notification subscription.")
	redshift_createEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.")
	redshift_createEventSubscriptionCmd.Flags().String("source-ids", "", "A list of one or more identifiers of Amazon Redshift source objects.")
	redshift_createEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that will be generating the events.")
	redshift_createEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the event subscription to be created.")
	redshift_createEventSubscriptionCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createEventSubscriptionCmd.MarkFlagRequired("sns-topic-arn")
	redshift_createEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	redshiftCmd.AddCommand(redshift_createEventSubscriptionCmd)
}
