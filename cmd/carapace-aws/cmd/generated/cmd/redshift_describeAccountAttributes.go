package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Returns a list of attributes attached to an account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeAccountAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeAccountAttributesCmd).Standalone()

		redshift_describeAccountAttributesCmd.Flags().String("attribute-names", "", "A list of attribute names.")
	})
	redshiftCmd.AddCommand(redshift_describeAccountAttributesCmd)
}
