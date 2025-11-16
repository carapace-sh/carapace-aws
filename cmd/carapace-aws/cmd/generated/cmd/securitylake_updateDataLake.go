package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_updateDataLakeCmd = &cobra.Command{
	Use:   "update-data-lake",
	Short: "You can use `UpdateDataLake` to specify where to store your security data, how it should be encrypted at rest and for how long.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_updateDataLakeCmd).Standalone()

	securitylake_updateDataLakeCmd.Flags().String("configurations", "", "Specifies the Region or Regions that will contribute data to the rollup region.")
	securitylake_updateDataLakeCmd.Flags().String("meta-store-manager-role-arn", "", "The Amazon Resource Name (ARN) used to create and update the Glue table.")
	securitylake_updateDataLakeCmd.MarkFlagRequired("configurations")
	securitylakeCmd.AddCommand(securitylake_updateDataLakeCmd)
}
