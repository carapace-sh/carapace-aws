package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteDataAccessorCmd = &cobra.Command{
	Use:   "delete-data-accessor",
	Short: "Deletes a specified data accessor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteDataAccessorCmd).Standalone()

	qbusiness_deleteDataAccessorCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
	qbusiness_deleteDataAccessorCmd.Flags().String("data-accessor-id", "", "The unique identifier of the data accessor to delete.")
	qbusiness_deleteDataAccessorCmd.MarkFlagRequired("application-id")
	qbusiness_deleteDataAccessorCmd.MarkFlagRequired("data-accessor-id")
	qbusinessCmd.AddCommand(qbusiness_deleteDataAccessorCmd)
}
