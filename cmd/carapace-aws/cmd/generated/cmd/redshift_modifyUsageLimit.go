package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyUsageLimitCmd = &cobra.Command{
	Use:   "modify-usage-limit",
	Short: "Modifies a usage limit in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyUsageLimitCmd).Standalone()

	redshift_modifyUsageLimitCmd.Flags().String("amount", "", "The new limit amount.")
	redshift_modifyUsageLimitCmd.Flags().String("breach-action", "", "The new action that Amazon Redshift takes when the limit is reached.")
	redshift_modifyUsageLimitCmd.Flags().String("usage-limit-id", "", "The identifier of the usage limit to modify.")
	redshift_modifyUsageLimitCmd.MarkFlagRequired("usage-limit-id")
	redshiftCmd.AddCommand(redshift_modifyUsageLimitCmd)
}
