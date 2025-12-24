package profile

import (
	"encoding/json"
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
)

const (
	EventProfileIndex = "profile.index"
)

func (p *ProfileService) sendProfileToIndex(profile model.Profile) {
	payload, err := json.Marshal(profile)
	if err != nil {
		log.Printf("failed to marshal profile: %v", err)
		return
	}

	if err := p.Bus.Publish(
		EventProfileIndex,
		[]byte(profile.Username),
		payload,
	); err != nil {
		log.Printf("failed to publish profile.index event: %v", err)
	}
}
