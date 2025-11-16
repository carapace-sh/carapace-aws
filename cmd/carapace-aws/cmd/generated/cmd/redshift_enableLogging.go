package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_enableLoggingCmd = &cobra.Command{
	Use:   "enable-logging",
	Short: "Starts logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_enableLoggingCmd).Standalone()

	redshift_enableLoggingCmd.Flags().String("bucket-name", "", "The name of an existing S3 bucket where the log files are to be stored.")
	redshift_enableLoggingCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster on which logging is to be started.")
	redshift_enableLoggingCmd.Flags().String("log-destination-type", "", "The log destination type.")
	redshift_enableLoggingCmd.Flags().String("log-exports", "", "The collection of exported log types.")
	redshift_enableLoggingCmd.Flags().String("s3-key-prefix", "", "The prefix applied to the log file names.")
	redshift_enableLoggingCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_enableLoggingCmd)
}
