package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeInsightCmd = &cobra.Command{
	Use:   "describe-insight",
	Short: "Returns details about an insight that you specify using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeInsightCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeInsightCmd).Standalone()

		devopsGuru_describeInsightCmd.Flags().String("account-id", "", "The ID of the member account in the organization.")
		devopsGuru_describeInsightCmd.Flags().String("id", "", "The ID of the insight.")
		devopsGuru_describeInsightCmd.MarkFlagRequired("id")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeInsightCmd)
}
