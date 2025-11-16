package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_rejectDataShareCmd = &cobra.Command{
	Use:   "reject-data-share",
	Short: "From a datashare consumer account, rejects the specified datashare.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_rejectDataShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_rejectDataShareCmd).Standalone()

		redshift_rejectDataShareCmd.Flags().String("data-share-arn", "", "The Amazon Resource Name (ARN) of the datashare to reject.")
		redshift_rejectDataShareCmd.MarkFlagRequired("data-share-arn")
	})
	redshiftCmd.AddCommand(redshift_rejectDataShareCmd)
}
