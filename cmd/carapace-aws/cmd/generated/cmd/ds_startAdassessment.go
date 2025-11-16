package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_startAdassessmentCmd = &cobra.Command{
	Use:   "start-adassessment",
	Short: "Initiates a directory assessment to validate your self-managed AD environment for hybrid domain join.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_startAdassessmentCmd).Standalone()

	ds_startAdassessmentCmd.Flags().String("assessment-configuration", "", "Configuration parameters for the directory assessment, including DNS server information, domain name, Amazon VPC subnet, and Amazon Web Services System Manager managed node details.")
	ds_startAdassessmentCmd.Flags().String("directory-id", "", "The identifier of the directory for which to perform the assessment.")
	dsCmd.AddCommand(ds_startAdassessmentCmd)
}
