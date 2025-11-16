package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_disassociateDataShareConsumerCmd = &cobra.Command{
	Use:   "disassociate-data-share-consumer",
	Short: "From a datashare consumer account, remove association for the specified datashare.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_disassociateDataShareConsumerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_disassociateDataShareConsumerCmd).Standalone()

		redshift_disassociateDataShareConsumerCmd.Flags().String("consumer-arn", "", "The Amazon Resource Name (ARN) of the consumer namespace that association for the datashare is removed from.")
		redshift_disassociateDataShareConsumerCmd.Flags().String("consumer-region", "", "From a datashare consumer account, removes association of a datashare from all the existing and future namespaces in the specified Amazon Web Services Region.")
		redshift_disassociateDataShareConsumerCmd.Flags().String("data-share-arn", "", "The Amazon Resource Name (ARN) of the datashare to remove association for.")
		redshift_disassociateDataShareConsumerCmd.Flags().String("disassociate-entire-account", "", "A value that specifies whether association for the datashare is removed from the entire account.")
		redshift_disassociateDataShareConsumerCmd.MarkFlagRequired("data-share-arn")
	})
	redshiftCmd.AddCommand(redshift_disassociateDataShareConsumerCmd)
}
