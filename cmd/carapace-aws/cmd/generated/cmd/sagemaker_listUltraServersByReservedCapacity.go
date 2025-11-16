package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listUltraServersByReservedCapacityCmd = &cobra.Command{
	Use:   "list-ultra-servers-by-reserved-capacity",
	Short: "Lists all UltraServers that are part of a specified reserved capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listUltraServersByReservedCapacityCmd).Standalone()

	sagemaker_listUltraServersByReservedCapacityCmd.Flags().String("max-results", "", "The maximum number of UltraServers to return in the response.")
	sagemaker_listUltraServersByReservedCapacityCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
	sagemaker_listUltraServersByReservedCapacityCmd.Flags().String("reserved-capacity-arn", "", "The ARN of the reserved capacity to list UltraServers for.")
	sagemaker_listUltraServersByReservedCapacityCmd.MarkFlagRequired("reserved-capacity-arn")
	sagemakerCmd.AddCommand(sagemaker_listUltraServersByReservedCapacityCmd)
}
