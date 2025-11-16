package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteUsageLimitCmd = &cobra.Command{
	Use:   "delete-usage-limit",
	Short: "Deletes a usage limit from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteUsageLimitCmd).Standalone()

	redshift_deleteUsageLimitCmd.Flags().String("usage-limit-id", "", "The identifier of the usage limit to delete.")
	redshift_deleteUsageLimitCmd.MarkFlagRequired("usage-limit-id")
	redshiftCmd.AddCommand(redshift_deleteUsageLimitCmd)
}
