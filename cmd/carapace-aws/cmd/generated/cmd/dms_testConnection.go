package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_testConnectionCmd = &cobra.Command{
	Use:   "test-connection",
	Short: "Tests the connection between the replication instance and the endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_testConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_testConnectionCmd).Standalone()

		dms_testConnectionCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
		dms_testConnectionCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
		dms_testConnectionCmd.MarkFlagRequired("endpoint-arn")
		dms_testConnectionCmd.MarkFlagRequired("replication-instance-arn")
	})
	dmsCmd.AddCommand(dms_testConnectionCmd)
}
