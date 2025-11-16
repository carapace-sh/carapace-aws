package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listMfadeviceTagsCmd = &cobra.Command{
	Use:   "list-mfadevice-tags",
	Short: "Lists the tags that are attached to the specified IAM virtual multi-factor authentication (MFA) device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listMfadeviceTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listMfadeviceTagsCmd).Standalone()

		iam_listMfadeviceTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listMfadeviceTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listMfadeviceTagsCmd.Flags().String("serial-number", "", "The unique identifier for the IAM virtual MFA device whose tags you want to see.")
		iam_listMfadeviceTagsCmd.MarkFlagRequired("serial-number")
	})
	iamCmd.AddCommand(iam_listMfadeviceTagsCmd)
}
