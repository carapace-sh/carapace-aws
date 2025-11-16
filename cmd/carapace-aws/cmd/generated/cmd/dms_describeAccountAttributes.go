package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Lists all of the DMS attributes for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeAccountAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeAccountAttributesCmd).Standalone()

	})
	dmsCmd.AddCommand(dms_describeAccountAttributesCmd)
}
