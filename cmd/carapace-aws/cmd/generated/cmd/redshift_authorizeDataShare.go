package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_authorizeDataShareCmd = &cobra.Command{
	Use:   "authorize-data-share",
	Short: "From a data producer account, authorizes the sharing of a datashare with one or more consumer accounts or managing entities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_authorizeDataShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_authorizeDataShareCmd).Standalone()

		redshift_authorizeDataShareCmd.Flags().String("allow-writes", "", "If set to true, allows write operations for a datashare.")
		redshift_authorizeDataShareCmd.Flags().String("consumer-identifier", "", "The identifier of the data consumer that is authorized to access the datashare.")
		redshift_authorizeDataShareCmd.Flags().String("data-share-arn", "", "The Amazon Resource Name (ARN) of the datashare namespace that producers are to authorize sharing for.")
		redshift_authorizeDataShareCmd.MarkFlagRequired("consumer-identifier")
		redshift_authorizeDataShareCmd.MarkFlagRequired("data-share-arn")
	})
	redshiftCmd.AddCommand(redshift_authorizeDataShareCmd)
}
