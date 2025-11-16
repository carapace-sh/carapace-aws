package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_associateDataShareConsumerCmd = &cobra.Command{
	Use:   "associate-data-share-consumer",
	Short: "From a datashare consumer account, associates a datashare with the account (AssociateEntireAccount) or the specified namespace (ConsumerArn).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_associateDataShareConsumerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_associateDataShareConsumerCmd).Standalone()

		redshift_associateDataShareConsumerCmd.Flags().String("allow-writes", "", "If set to true, allows write operations for a datashare.")
		redshift_associateDataShareConsumerCmd.Flags().String("associate-entire-account", "", "A value that specifies whether the datashare is associated with the entire account.")
		redshift_associateDataShareConsumerCmd.Flags().String("consumer-arn", "", "The Amazon Resource Name (ARN) of the consumer namespace associated with the datashare.")
		redshift_associateDataShareConsumerCmd.Flags().String("consumer-region", "", "From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified Amazon Web Services Region.")
		redshift_associateDataShareConsumerCmd.Flags().String("data-share-arn", "", "The Amazon Resource Name (ARN) of the datashare that the consumer is to use.")
		redshift_associateDataShareConsumerCmd.MarkFlagRequired("data-share-arn")
	})
	redshiftCmd.AddCommand(redshift_associateDataShareConsumerCmd)
}
