package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeLoggingStatusCmd = &cobra.Command{
	Use:   "describe-logging-status",
	Short: "Describes whether information, such as queries and connection attempts, is being logged for the specified Amazon Redshift cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeLoggingStatusCmd).Standalone()

	redshift_describeLoggingStatusCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster from which to get the logging status.")
	redshift_describeLoggingStatusCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_describeLoggingStatusCmd)
}
