package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteAccountPoolCmd = &cobra.Command{
	Use:   "delete-account-pool",
	Short: "Deletes an account pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteAccountPoolCmd).Standalone()

	datazone_deleteAccountPoolCmd.Flags().String("domain-identifier", "", "The ID of the domain where the account pool is deleted.")
	datazone_deleteAccountPoolCmd.Flags().String("identifier", "", "The ID of the account pool to be deleted.")
	datazone_deleteAccountPoolCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteAccountPoolCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteAccountPoolCmd)
}
