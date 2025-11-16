package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Lists all of the attributes for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeAccountAttributesCmd).Standalone()

	rdsCmd.AddCommand(rds_describeAccountAttributesCmd)
}
