package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeHsmConfigurationsCmd = &cobra.Command{
	Use:   "describe-hsm-configurations",
	Short: "Returns information about the specified Amazon Redshift HSM configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeHsmConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeHsmConfigurationsCmd).Standalone()

		redshift_describeHsmConfigurationsCmd.Flags().String("hsm-configuration-identifier", "", "The identifier of a specific Amazon Redshift HSM configuration to be described.")
		redshift_describeHsmConfigurationsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeHsmConfigurationsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeHsmConfigurationsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching HSM configurations that are associated with the specified key or keys.")
		redshift_describeHsmConfigurationsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching HSM configurations that are associated with the specified tag value or values.")
	})
	redshiftCmd.AddCommand(redshift_describeHsmConfigurationsCmd)
}
