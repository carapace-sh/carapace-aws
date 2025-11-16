package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getDataAccessorCmd = &cobra.Command{
	Use:   "get-data-accessor",
	Short: "Retrieves information about a specified data accessor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getDataAccessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getDataAccessorCmd).Standalone()

		qbusiness_getDataAccessorCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
		qbusiness_getDataAccessorCmd.Flags().String("data-accessor-id", "", "The unique identifier of the data accessor to retrieve.")
		qbusiness_getDataAccessorCmd.MarkFlagRequired("application-id")
		qbusiness_getDataAccessorCmd.MarkFlagRequired("data-accessor-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getDataAccessorCmd)
}
