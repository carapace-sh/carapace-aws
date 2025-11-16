package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getRegisterAccountStatusCmd = &cobra.Command{
	Use:   "get-register-account-status",
	Short: "Retrieves information about the status of registering your Amazon Web Services account, IAM, and Amazon Timestream resources so that Amazon Web Services IoT FleetWise can transfer your vehicle data to the Amazon Web Services Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getRegisterAccountStatusCmd).Standalone()

	iotfleetwiseCmd.AddCommand(iotfleetwise_getRegisterAccountStatusCmd)
}
