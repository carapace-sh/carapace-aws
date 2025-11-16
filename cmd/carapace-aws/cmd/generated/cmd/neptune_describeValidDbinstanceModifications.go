package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeValidDbinstanceModificationsCmd = &cobra.Command{
	Use:   "describe-valid-dbinstance-modifications",
	Short: "You can call [DescribeValidDBInstanceModifications]() to learn what modifications you can make to your DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeValidDbinstanceModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeValidDbinstanceModificationsCmd).Standalone()

		neptune_describeValidDbinstanceModificationsCmd.Flags().String("dbinstance-identifier", "", "The customer identifier or the ARN of your DB instance.")
		neptune_describeValidDbinstanceModificationsCmd.MarkFlagRequired("dbinstance-identifier")
	})
	neptuneCmd.AddCommand(neptune_describeValidDbinstanceModificationsCmd)
}
