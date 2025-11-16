package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteInstanceProfileCmd = &cobra.Command{
	Use:   "delete-instance-profile",
	Short: "Deletes the specified instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteInstanceProfileCmd).Standalone()

		dms_deleteInstanceProfileCmd.Flags().String("instance-profile-identifier", "", "The identifier of the instance profile to delete.")
		dms_deleteInstanceProfileCmd.MarkFlagRequired("instance-profile-identifier")
	})
	dmsCmd.AddCommand(dms_deleteInstanceProfileCmd)
}
