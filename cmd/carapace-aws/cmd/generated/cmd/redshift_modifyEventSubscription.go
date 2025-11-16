package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyEventSubscriptionCmd = &cobra.Command{
	Use:   "modify-event-subscription",
	Short: "Modifies an existing Amazon Redshift event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyEventSubscriptionCmd).Standalone()

	redshift_modifyEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value indicating if the subscription is enabled.")
	redshift_modifyEventSubscriptionCmd.Flags().String("event-categories", "", "Specifies the Amazon Redshift event categories to be published by the event notification subscription.")
	redshift_modifyEventSubscriptionCmd.Flags().String("severity", "", "Specifies the Amazon Redshift event severity to be published by the event notification subscription.")
	redshift_modifyEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic to be used by the event notification subscription.")
	redshift_modifyEventSubscriptionCmd.Flags().String("source-ids", "", "A list of one or more identifiers of Amazon Redshift source objects.")
	redshift_modifyEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that will be generating the events.")
	redshift_modifyEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the modified Amazon Redshift event notification subscription.")
	redshift_modifyEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	redshiftCmd.AddCommand(redshift_modifyEventSubscriptionCmd)
}
