package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeAnomalyCmd = &cobra.Command{
	Use:   "describe-anomaly",
	Short: "Returns details about an anomaly that you specify using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeAnomalyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeAnomalyCmd).Standalone()

		devopsGuru_describeAnomalyCmd.Flags().String("account-id", "", "The ID of the member account.")
		devopsGuru_describeAnomalyCmd.Flags().String("id", "", "The ID of the anomaly.")
		devopsGuru_describeAnomalyCmd.MarkFlagRequired("id")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeAnomalyCmd)
}
