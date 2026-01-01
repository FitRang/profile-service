package proto

import "github.com/Foxtrot-14/FitRang/profile-service/graph/model"

func toProtoProfile(p model.Profile) *Profile {
	var profileURL string
	if p.ProfileURL != nil {
		profileURL = *p.ProfileURL
	}

	return &Profile{
		Id:           p.ID,
		FullName:     p.FullName,
		Username:     p.Username,
		ProfileUrl:   profileURL,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func toProtoProfiles(profiles []model.Profile) []*Profile {
	out := make([]*Profile, 0, len(profiles))
	for _, p := range profiles {
		out = append(out, toProtoProfile(p))
	}
	return out
}

func toProtoDossier(d model.Dossier) *Dossier {
	var height string
	if d.Height != nil {
		height = *d.Height
	}

	var weight string
	if d.Weight != nil {
		weight = *d.Weight
	}

	preferredColors := []string{}
	if d.PreferredColors != nil {
		preferredColors = d.PreferredColors
	}

	dislikedColors := []string{}
	if d.DislikedColors != nil {
		dislikedColors = d.DislikedColors
	}

	viewers := []string{}
	if d.Viewers != nil {
		viewers = d.Viewers
	}

	return &Dossier{
		Id:              d.ID,
		Username:        d.Username,
		FaceType:        d.FaceType,
		SkinTone:        d.SkinTone.String(),
		BodyType:        d.BodyType,
		Gender:          d.Gender.String(),
		PreferredColors: preferredColors,
		DislikedColors:  dislikedColors,
		Viewers:         viewers,
		Height:          height,
		Weight:          weight,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}
