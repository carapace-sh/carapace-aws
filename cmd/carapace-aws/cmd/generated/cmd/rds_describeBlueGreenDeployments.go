package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeBlueGreenDeploymentsCmd = &cobra.Command{
	Use:   "describe-blue-green-deployments",
	Short: "Describes one or more blue/green deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeBlueGreenDeploymentsCmd).Standalone()

	rds_describeBlueGreenDeploymentsCmd.Flags().String("blue-green-deployment-identifier", "", "The blue/green deployment identifier.")
	rds_describeBlueGreenDeploymentsCmd.Flags().String("filters", "", "A filter that specifies one or more blue/green deployments to describe.")
	rds_describeBlueGreenDeploymentsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeBlueGreenDeployments` request.")
	rds_describeBlueGreenDeploymentsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeBlueGreenDeploymentsCmd)
}
