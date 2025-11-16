package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createDataAccessorCmd = &cobra.Command{
	Use:   "create-data-accessor",
	Short: "Creates a new data accessor for an ISV to access data from a Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createDataAccessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createDataAccessorCmd).Standalone()

		qbusiness_createDataAccessorCmd.Flags().String("action-configurations", "", "A list of action configurations specifying the allowed actions and any associated filters.")
		qbusiness_createDataAccessorCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
		qbusiness_createDataAccessorCmd.Flags().String("authentication-detail", "", "The authentication configuration details for the data accessor.")
		qbusiness_createDataAccessorCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		qbusiness_createDataAccessorCmd.Flags().String("display-name", "", "A friendly name for the data accessor.")
		qbusiness_createDataAccessorCmd.Flags().String("principal", "", "The Amazon Resource Name (ARN) of the IAM role for the ISV that will be accessing the data.")
		qbusiness_createDataAccessorCmd.Flags().String("tags", "", "The tags to associate with the data accessor.")
		qbusiness_createDataAccessorCmd.MarkFlagRequired("action-configurations")
		qbusiness_createDataAccessorCmd.MarkFlagRequired("application-id")
		qbusiness_createDataAccessorCmd.MarkFlagRequired("display-name")
		qbusiness_createDataAccessorCmd.MarkFlagRequired("principal")
	})
	qbusinessCmd.AddCommand(qbusiness_createDataAccessorCmd)
}
