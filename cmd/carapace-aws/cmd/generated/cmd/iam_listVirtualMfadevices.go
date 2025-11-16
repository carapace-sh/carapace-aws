package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listVirtualMfadevicesCmd = &cobra.Command{
	Use:   "list-virtual-mfadevices",
	Short: "Lists the virtual MFA devices defined in the Amazon Web Services account by assignment status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listVirtualMfadevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listVirtualMfadevicesCmd).Standalone()

		iam_listVirtualMfadevicesCmd.Flags().String("assignment-status", "", "The status (`Unassigned` or `Assigned`) of the devices to list.")
		iam_listVirtualMfadevicesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listVirtualMfadevicesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	})
	iamCmd.AddCommand(iam_listVirtualMfadevicesCmd)
}
