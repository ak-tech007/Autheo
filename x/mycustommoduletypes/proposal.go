// mycustommoduletypes/proposal.go
package mycustommoduletypes

import (
	"fmt"

	"Autheo/x/gov/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type MyCustomProposal struct {
	Title       string
	Description string
	// Add any additional fields here...
}

var _ govtypes.Content = MyCustomProposal{}

func (mcp MyCustomProposal) GetTitle() string {
	return "My Custom Proposal Title"
}

func (mcp MyCustomProposal) GetDescription() string {
	return "Description of My Custom Proposal."
}

func (mcp MyCustomProposal) String() string {
	// Return a string representation of your proposal
	return fmt.Sprintf("Title: %s\nDescription: %s", mcp.Title, mcp.Description)
}

func (mcp MyCustomProposal) ProposalRoute() string {
	return types.RouterKey // The router key for your module
}

func (mcp MyCustomProposal) ProposalType() string {
	return "MyCustomProposal" // Unique type identifier
}

func (mcp MyCustomProposal) ValidateBasic() error {
	if mcp.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if mcp.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return nil
}
