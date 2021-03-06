// Package camera provides data structures and adheres to the AvvaOscMessageBuilder interface
package camera

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hypebeast/go-osc/osc"
)

type Camera struct {
	ClearFlags  `json:",omitempty"`
	FieldOfView FieldOfView `json:",omitempty"`
}

// Generate generates an OSC message which can be sent to the avva.studio visuals system.
// Generate method makes Camera type adhere to AvvaOscMessageBuilder interface
func (c Camera) Generate() *osc.Message {
	json, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error creating camera json: %s", err.Error())
	}
	return osc.NewMessage("/camera", string(json))
}

// String makes Canvas adhere to the Stringer interface
func (c Camera) String() string {
	return fmt.Sprintf("%s %s", "camera", c.ClearFlags)
}
