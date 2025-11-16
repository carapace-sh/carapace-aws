package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getInstanceProfileCmd = &cobra.Command{
	Use:   "get-instance-profile",
	Short: "Returns information about the specified instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getInstanceProfileCmd).Standalone()

	devicefarm_getInstanceProfileCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of an instance profile.")
	devicefarm_getInstanceProfileCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getInstanceProfileCmd)
}
