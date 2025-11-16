package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createDataLakeCmd = &cobra.Command{
	Use:   "create-data-lake",
	Short: "Initializes an Amazon Security Lake instance with the provided (or default) configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createDataLakeCmd).Standalone()

	securitylake_createDataLakeCmd.Flags().String("configurations", "", "Specify the Region or Regions that will contribute data to the rollup region.")
	securitylake_createDataLakeCmd.Flags().String("meta-store-manager-role-arn", "", "The Amazon Resource Name (ARN) used to create and update the Glue table.")
	securitylake_createDataLakeCmd.Flags().String("tags", "", "An array of objects, one for each tag to associate with the data lake configuration.")
	securitylake_createDataLakeCmd.MarkFlagRequired("configurations")
	securitylake_createDataLakeCmd.MarkFlagRequired("meta-store-manager-role-arn")
	securitylakeCmd.AddCommand(securitylake_createDataLakeCmd)
}
