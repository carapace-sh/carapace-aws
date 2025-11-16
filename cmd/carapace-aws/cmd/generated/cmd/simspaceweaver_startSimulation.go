package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_startSimulationCmd = &cobra.Command{
	Use:   "start-simulation",
	Short: "Starts a simulation with the given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_startSimulationCmd).Standalone()

	simspaceweaver_startSimulationCmd.Flags().String("client-token", "", "A value that you provide to ensure that repeated calls to this API operation using the same parameters complete only once.")
	simspaceweaver_startSimulationCmd.Flags().String("description", "", "The description of the simulation.")
	simspaceweaver_startSimulationCmd.Flags().String("maximum-duration", "", "The maximum running time of the simulation, specified as a number of minutes (m or M), hours (h or H), or days (d or D).")
	simspaceweaver_startSimulationCmd.Flags().String("name", "", "The name of the simulation.")
	simspaceweaver_startSimulationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the simulation assumes to perform actions.")
	simspaceweaver_startSimulationCmd.Flags().String("schema-s3-location", "", "The location of the simulation schema in Amazon Simple Storage Service (Amazon S3).")
	simspaceweaver_startSimulationCmd.Flags().String("snapshot-s3-location", "", "The location of the snapshot .zip file in Amazon Simple Storage Service (Amazon S3).")
	simspaceweaver_startSimulationCmd.Flags().String("tags", "", "A list of tags for the simulation.")
	simspaceweaver_startSimulationCmd.MarkFlagRequired("name")
	simspaceweaver_startSimulationCmd.MarkFlagRequired("role-arn")
	simspaceweaverCmd.AddCommand(simspaceweaver_startSimulationCmd)
}
