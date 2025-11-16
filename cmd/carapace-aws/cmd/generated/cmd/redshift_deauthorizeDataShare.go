package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deauthorizeDataShareCmd = &cobra.Command{
	Use:   "deauthorize-data-share",
	Short: "From a datashare producer account, removes authorization from the specified datashare.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deauthorizeDataShareCmd).Standalone()

	redshift_deauthorizeDataShareCmd.Flags().String("consumer-identifier", "", "The identifier of the data consumer that is to have authorization removed from the datashare.")
	redshift_deauthorizeDataShareCmd.Flags().String("data-share-arn", "", "The namespace Amazon Resource Name (ARN) of the datashare to remove authorization from.")
	redshift_deauthorizeDataShareCmd.MarkFlagRequired("consumer-identifier")
	redshift_deauthorizeDataShareCmd.MarkFlagRequired("data-share-arn")
	redshiftCmd.AddCommand(redshift_deauthorizeDataShareCmd)
}
