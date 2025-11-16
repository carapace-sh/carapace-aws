package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getInstanceProfileCmd = &cobra.Command{
	Use:   "get-instance-profile",
	Short: "Retrieves information about the specified instance profile, including the instance profile's path, GUID, ARN, and role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getInstanceProfileCmd).Standalone()

	iam_getInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the instance profile to get information about.")
	iam_getInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
	iamCmd.AddCommand(iam_getInstanceProfileCmd)
}
