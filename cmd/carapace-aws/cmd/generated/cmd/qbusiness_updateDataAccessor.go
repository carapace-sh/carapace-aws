package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateDataAccessorCmd = &cobra.Command{
	Use:   "update-data-accessor",
	Short: "Updates an existing data accessor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateDataAccessorCmd).Standalone()

	qbusiness_updateDataAccessorCmd.Flags().String("action-configurations", "", "The updated list of action configurations specifying the allowed actions and any associated filters.")
	qbusiness_updateDataAccessorCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
	qbusiness_updateDataAccessorCmd.Flags().String("authentication-detail", "", "The updated authentication configuration details for the data accessor.")
	qbusiness_updateDataAccessorCmd.Flags().String("data-accessor-id", "", "The unique identifier of the data accessor to update.")
	qbusiness_updateDataAccessorCmd.Flags().String("display-name", "", "The updated friendly name for the data accessor.")
	qbusiness_updateDataAccessorCmd.MarkFlagRequired("action-configurations")
	qbusiness_updateDataAccessorCmd.MarkFlagRequired("application-id")
	qbusiness_updateDataAccessorCmd.MarkFlagRequired("data-accessor-id")
	qbusinessCmd.AddCommand(qbusiness_updateDataAccessorCmd)
}
