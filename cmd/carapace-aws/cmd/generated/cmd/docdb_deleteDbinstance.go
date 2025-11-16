package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteDbinstanceCmd = &cobra.Command{
	Use:   "delete-dbinstance",
	Short: "Deletes a previously provisioned instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteDbinstanceCmd).Standalone()

	docdb_deleteDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The instance identifier for the instance to be deleted.")
	docdb_deleteDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	docdbCmd.AddCommand(docdb_deleteDbinstanceCmd)
}
