package main

import (
	"github.com/cloudfoundry/bosh-cli/ui/table"
)

type WriterUI struct {
	Block []byte
}

func (ui *WriterUI) IsTTY() bool {
	return false
}

// ErrorLinef starts and ends a text error line
func (ui *WriterUI) ErrorLinef(pattern string, args ...interface{}) {
	panic("Not Implmented")
}

// Printlnf starts and ends a text line
func (ui *WriterUI) PrintLinef(pattern string, args ...interface{}) {
	panic("Not Implmented")
}

// PrintBeginf starts a text line
func (ui *WriterUI) BeginLinef(pattern string, args ...interface{}) {
	panic("Not Implmented")
}

// PrintEndf ends a text line
func (ui *WriterUI) EndLinef(pattern string, args ...interface{}) {
	panic("Not Implmented")
}

func (ui *WriterUI) PrintBlock(block []byte) {
	ui.Block = append(ui.Block, block...)
}

func (ui *WriterUI) PrintErrorBlock(block string) {
	panic("Not Implmented")
}

func (ui *WriterUI) PrintTable(table table.Table) {
	panic("Not Implmented")
}

func (ui *WriterUI) AskForText(label string) (string, error) {
	panic("Not Implmented")
}

func (ui *WriterUI) AskForChoice(label string, options []string) (int, error) {
	panic("Not Implmented")
}

func (ui *WriterUI) AskForPassword(label string) (string, error) {
	panic("Not Implmented")
}

func (ui *WriterUI) AskForConfirmation() error {
	panic("Not Implmented")
}

func (ui *WriterUI) IsInteractive() bool {
	return true
}

func (ui *WriterUI) Flush() {}
