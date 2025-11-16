package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAssociationCmd = &cobra.Command{
	Use:   "describe-association",
	Short: "Describes the association for the specified target or managed node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAssociationCmd).Standalone()

	ssm_describeAssociationCmd.Flags().String("association-id", "", "The association ID for which you want information.")
	ssm_describeAssociationCmd.Flags().String("association-version", "", "Specify the association version to retrieve.")
	ssm_describeAssociationCmd.Flags().String("instance-id", "", "The managed node ID.")
	ssm_describeAssociationCmd.Flags().String("name", "", "The name of the SSM document.")
	ssmCmd.AddCommand(ssm_describeAssociationCmd)
}
