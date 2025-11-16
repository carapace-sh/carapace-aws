package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_registerAccountCmd = &cobra.Command{
	Use:   "register-account",
	Short: "This API operation contains deprecated parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_registerAccountCmd).Standalone()

	iotfleetwise_registerAccountCmd.Flags().String("iam-resources", "", "The IAM resource that allows Amazon Web Services IoT FleetWise to send data to Amazon Timestream.")
	iotfleetwise_registerAccountCmd.Flags().String("timestream-resources", "", "")
	iotfleetwiseCmd.AddCommand(iotfleetwise_registerAccountCmd)
}
