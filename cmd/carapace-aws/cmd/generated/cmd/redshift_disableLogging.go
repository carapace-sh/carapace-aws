package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_disableLoggingCmd = &cobra.Command{
	Use:   "disable-logging",
	Short: "Stops logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_disableLoggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_disableLoggingCmd).Standalone()

		redshift_disableLoggingCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster on which logging is to be stopped.")
		redshift_disableLoggingCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_disableLoggingCmd)
}
