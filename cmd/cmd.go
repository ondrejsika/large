package cmd

import (
	_ "github.com/ondrejsika/training-cli/cmd"
	training_cli_root "github.com/ondrejsika/training-cli/cmd/root"
	_ "github.com/sikalabs/signpost/cmd"
	signpost_root "github.com/sikalabs/signpost/cmd/root"
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var LargeCmd = &cobra.Command{
	Use:   "large",
	Short: "large: slu, signpost & training-cli in large binary",
}

func init() {
	LargeCmd.AddCommand(slu_root.RootCmd)
	LargeCmd.AddCommand(training_cli_root.Cmd)
	LargeCmd.AddCommand(signpost_root.Cmd)
}

func Execute() {
	LargeCmd.Execute()
}
