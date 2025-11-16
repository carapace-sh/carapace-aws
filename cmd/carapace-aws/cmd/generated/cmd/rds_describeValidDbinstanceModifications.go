package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeValidDbinstanceModificationsCmd = &cobra.Command{
	Use:   "describe-valid-dbinstance-modifications",
	Short: "You can call `DescribeValidDBInstanceModifications` to learn what modifications you can make to your DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeValidDbinstanceModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeValidDbinstanceModificationsCmd).Standalone()

		rds_describeValidDbinstanceModificationsCmd.Flags().String("dbinstance-identifier", "", "The customer identifier or the ARN of your DB instance.")
		rds_describeValidDbinstanceModificationsCmd.MarkFlagRequired("dbinstance-identifier")
	})
	rdsCmd.AddCommand(rds_describeValidDbinstanceModificationsCmd)
}
