package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_createExtendedSourceServerCmd = &cobra.Command{
	Use:   "create-extended-source-server",
	Short: "Create an extended source server in the target Account based on the source server in staging account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_createExtendedSourceServerCmd).Standalone()

	drs_createExtendedSourceServerCmd.Flags().String("source-server-arn", "", "This defines the ARN of the source server in staging Account based on which you want to create an extended source server.")
	drs_createExtendedSourceServerCmd.Flags().String("tags", "", "A list of tags associated with the extended source server.")
	drs_createExtendedSourceServerCmd.MarkFlagRequired("source-server-arn")
	drsCmd.AddCommand(drs_createExtendedSourceServerCmd)
}
