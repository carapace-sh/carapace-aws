package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteInstanceProfileCmd = &cobra.Command{
	Use:   "delete-instance-profile",
	Short: "Deletes a profile that can be applied to one or more private device instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteInstanceProfileCmd).Standalone()

	devicefarm_deleteInstanceProfileCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the instance profile you are requesting to delete.")
	devicefarm_deleteInstanceProfileCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_deleteInstanceProfileCmd)
}
