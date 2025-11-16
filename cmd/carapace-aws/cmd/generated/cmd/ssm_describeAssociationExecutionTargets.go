package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAssociationExecutionTargetsCmd = &cobra.Command{
	Use:   "describe-association-execution-targets",
	Short: "Views information about a specific execution of a specific association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAssociationExecutionTargetsCmd).Standalone()

	ssm_describeAssociationExecutionTargetsCmd.Flags().String("association-id", "", "The association ID that includes the execution for which you want to view details.")
	ssm_describeAssociationExecutionTargetsCmd.Flags().String("execution-id", "", "The execution ID for which you want to view details.")
	ssm_describeAssociationExecutionTargetsCmd.Flags().String("filters", "", "Filters for the request.")
	ssm_describeAssociationExecutionTargetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeAssociationExecutionTargetsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_describeAssociationExecutionTargetsCmd.MarkFlagRequired("association-id")
	ssm_describeAssociationExecutionTargetsCmd.MarkFlagRequired("execution-id")
	ssmCmd.AddCommand(ssm_describeAssociationExecutionTargetsCmd)
}
