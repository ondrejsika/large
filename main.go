package main

import (
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

var LargeCmd = &cobra.Command{
	Use:   "large",
	Short: "slu, tergum, ... in large binary",
}

func init() {
	LargeCmd.AddCommand(slu_root.RootCmd)
	LargeCmd.AddCommand(tergum_root.Cmd)
}

func main() {
	LargeCmd.Execute()
}
