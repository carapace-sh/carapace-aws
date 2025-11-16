package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes the connection between a replication instance and an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteConnectionCmd).Standalone()

	dms_deleteConnectionCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
	dms_deleteConnectionCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	dms_deleteConnectionCmd.MarkFlagRequired("endpoint-arn")
	dms_deleteConnectionCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_deleteConnectionCmd)
}
